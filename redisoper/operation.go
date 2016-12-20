package redisoper

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

type redisServer struct {
	redisHost string
	redisAuth string
}

//get redisServer entity
//example redisHost "127.0.0.1:6379" redisPassword "123"
func NewRedis(redisHost, redisAuth string) *redisServer {
	return &redisServer{redisHost, redisAuth}
}

func (rs *redisServer) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", rs.redisHost)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", rs.redisAuth); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (_ *redisServer) GetData(pool *redis.Pool, key, keyType string, args ...interface{}) (data interface{}, err error) {
	conn := pool.Get()
	defer conn.Close()

	switch keyType {
	case "string":
		data, err = redis.Bytes(conn.Do("get", key))
	case "list":
		argArrList := make([]interface{}, 2)
		if len(args) > 2 {
			argArrList = args[0:2]
		} else {
			argArrList = args
		}
		arrList := make([]int, 2)
		for index, value := range argArrList {
			if _, ok := value.(int); ok {
				arrList[index] = value.(int)
			} else {
				data, err = nil, errors.New("list range error")
				return
			}
		}
		data, err = redis.Strings(conn.Do("lrange", key, arrList[0], arrList[1]))
	case "hash":
		argArrHash := make([]interface{}, 1)
		if len(args) > 1 {
			argArrHash = args[0:1]
		} else {
			argArrHash = args
		}
		arrHash := make([]string, 1)
		for index, value := range argArrHash {
			if _, ok := value.(string); ok {
				arrHash[index] = value.(string)
			} else {
				data, err = nil, errors.New("hash field error")
				return
			}
		}
		data, err = redis.String(conn.Do("hget", key, arrHash[0]))
	case "set":
		data, err = redis.Strings(conn.Do("smembers", key))
	case "sortset":
		argArrSort := make([]interface{}, 1)
		arrSort := make([]int, 2)

		if len(args) > 2 {
			argArrSort = args[0:2]
		} else {
			argArrSort = args
		}
		for index, value := range argArrSort {
			if _, ok := value.(int); ok {
				arrSort[index] = value.(int)
			} else {
				data, err = nil, errors.New("set range error")
				return
			}
		}
		data, err = redis.Strings(conn.Do("zrange", key, arrSort[0], arrSort[1]))
	default:
		data, err = nil, errors.New("input type err")
	}
	return
}
