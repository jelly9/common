package main

import (
	"fmt"
	"common/redis"
)

func main() {
	opt := &redis.Options{
		Addr: "39.105.149.213:6379",
	}
	client := redis.NewRedisClient(opt)
	//client.Set("aaa", "10")
	value, _ := client.Get("bbb").Result()
	fmt.Println("value: ", value)
}