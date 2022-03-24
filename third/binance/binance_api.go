package binance

import (
	"binance_project/global"
	"context"
	"encoding/json"
	"log"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	test_api  = "wss://stream.yshyqxx.com/stream"
	store_key = "streams=btcusdt@aggTrade"
)

func Get_Collection_Transactions() {
	ctx := context.Background()
	w, _, err := websocket.Dial(ctx, test_api+"?streams=btcusdt@aggTrade", nil)

	if err != nil {
		log.Printf("Websocket err %s", err.Error())
		defer w.Close(websocket.StatusInternalError, err.Error())
	}

	var v interface{}

	for {
		err = wsjson.Read(ctx, w, &v)
		apiData, err := json.Marshal(v)

		if err != nil {
			log.Printf("apiData Turn to byte error: %s", err.Error())
			defer w.Close(websocket.StatusInternalError, err.Error())
		}

		err = global.RedisDB.Set(ctx, store_key, apiData, 0).Err()

		if err != nil {
			log.Printf("Set Redis Key Err: %s", err.Error())
			defer w.Close(websocket.StatusInternalError, err.Error())
		}
	}

}
