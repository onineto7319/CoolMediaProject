package database

import "github.com/go-redis/redis/v8"

func NewRedisEngine() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb, nil
}
