package main

import (
	"fmt"

	"./redisoper"
)

func main() {
	redis := redisoper.NewRedis("127.0.0.1:6379", "123")
	pool := redis.NewPool()
	res, err := redis.GetDate(pool, "t2", "set")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.GetDate(pool, "t1", "string")
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(res)
	}
}
