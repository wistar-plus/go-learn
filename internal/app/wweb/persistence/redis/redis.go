package redis

import (
	"github.com/go-redis/redis/v7"
)

var db *redis.Client

func Init() {
	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Close() {
	if db != nil {
		db.Close()
	}
}
