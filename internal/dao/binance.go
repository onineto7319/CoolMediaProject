package dao

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	binanceInstance *binance
)

type binance struct {
	rdb *redis.Client
}

type binacne_interface interface {
	Get(key string) (interface{}, error)
}

func NewBinance(rdb *redis.Client) binacne_interface {
	var once sync.Once

	if binanceInstance == nil {
		once.Do(func() { binanceInstance = &binance{rdb: rdb} })
	}
	return binanceInstance
}

func (b *binance) Get(key string) (interface{}, error) {
	val, err := b.rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Get Redis Value Error: %s", err.Error())
		return nil, err
	}

	return val, nil
}
