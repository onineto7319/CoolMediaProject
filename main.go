package main

import (
	"binance_project/global"
	"binance_project/internal/router"
	"binance_project/pkg/database"
	"binance_project/third/binance"
	"log"
	"net/http"
)

func init() {
	err := setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	r := router.New()

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go binance.Get_Collection_Transactions()

	log.Fatal(s.ListenAndServe())
}

func setupDBEngine() error {
	var err error

	global.RedisDB, err = database.NewRedisEngine()

	if err != nil {
		return err
	}

	return err
}
