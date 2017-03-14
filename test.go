package main

import (
	"fmt"

	"./redisoper"
)

func main() {
	redis := redisoper.NewRedis("192.168.1.91:6379", "123")
	res, err := redis.Get("t1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res.([]byte)))
	}
	res, err = redis.Set("t3", "22")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.Sadd("s1", "2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
