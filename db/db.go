package db

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func Init() *redis.Client {
	// new redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})

	// test connection
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	// return pong if server is online
	fmt.Println(pong)

	return client
}
