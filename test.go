package main

import (
	"fmt"

	"./redisoper"
)

func main() {
	redis := redisoper.NewRedis("127.0.0.1:6379", "123")
	pool := redis.NewPool()

	res, err := redis.WriteData(pool, "zzz", "test1", "sortset", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.GetData(pool, "zzz", "sortset", 0, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// res, err := redis.GetData(pool, "t2", "set")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	// res, err = redis.GetData(pool, "t1", "string")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	// res, err = redis.GetData(pool, "l1", "list", 0, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	// res, err = redis.GetData(pool, "h1", "hash", "t1")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {

	// 	fmt.Println(res)
	// }

	// res, err = redis.GetData(pool, "z1", "sortset", 0, 3)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }
}
