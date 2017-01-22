package main

import (
	"fmt"

	"./redisoper"
)

func main() {
	redis := redisoper.NewRedis("192.168.1.41:6379", "123")
	res, err := redis.WriteData("zzz", "test2", "sortset", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.GetData("zzz", "sortset", 0, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.DelData("zzz")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = redis.WriteData("xxx", "ssss", "set")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	res, err = redis.WriteData("xxx", "sss", "set")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	res, err = redis.RemData("xxx", "ssss", "set")
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
