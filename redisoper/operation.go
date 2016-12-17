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

func (_ *redisServer) GetDate(pool *redis.Pool, key, keyType string) (data interface{}, err error) {
	conn := pool.Get()
	defer conn.Close()

	switch keyType {
	case "string":
		data, err = redis.Bytes(conn.Do("get", key))
	case "list":
		//data, err = redis.Strings(conn.Do("lpop", key))
	case "hash":

	case "set":
		data, err = redis.Strings(conn.Do("smembers", key))
	case "sortset":
		data, err = redis.Strings(conn.Do("zrange", key))
	default:
		data, err = nil, errors.New("input type err")
	}
	return
}
