package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil {
		log.Fatal("Filed to connect to Redis:", err)
	}

	fmt.Println("Redis connected successfully")
}
