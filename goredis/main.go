package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println("ping result", pong, err)

	err = client.Set("key1", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	val2, err := client.Get("key3").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}
