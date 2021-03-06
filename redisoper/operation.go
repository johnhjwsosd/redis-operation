package redisoper

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//RedisServer 通过NewRedis得到对象
type RedisServer struct {
	pool *redis.Pool
}

//NewRedis get redisServer entity
//returns 返回一个...
//example redisHost "127.0.0.1:6379" redisPassword "123" maxIdle 1000 maxActive 1000
func NewRedis(redisHost, redisAuth string,maxIdle,maxActive int) *RedisServer {
	return &RedisServer{newPool(redisHost, redisAuth,maxIdle,maxActive)}
}

func newPool(redisHost, redisAuth string,maxIdle,maxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", redisAuth); err != nil {
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

func (r *RedisServer) Ping()error{
	conn:=r.pool.Get()
	defer conn.Close()
	err := r.pool.TestOnBorrow(conn,time.Now().Add(-time.Minute))
	return err
}

//DelData ...
func (r *RedisServer) DelData(key string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("del", key))
	return
}

func (r *RedisServer) Incr(key string)(data interface{},err error){
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("incr", key))
	return
}

func (r *RedisServer) IncrBY(key string,value int)(data interface{},err error){
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("incrby", key,value))
	return
}


//Get String get
func (r *RedisServer) Get(key string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Bytes(conn.Do("get", key))
	return
}

//Set String set
func (r *RedisServer) Set(key string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("set", key, value))
	return
}

func (r *RedisServer) SetEx(key string,exp int,value interface{})(data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("setex", key,exp,value))
	return
}

//Lrange ...
func (r *RedisServer) Lrange(key string, startIndex, endIndex int) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Strings(conn.Do("lrange", key, startIndex, endIndex))
	return
}

//Lpush ...
func (r *RedisServer) Lpush(key string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("lpush", key, value))
	return
}

//Hget ...
func (r *RedisServer) Hget(key, field string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("hget", key, field))
	return
}

//Hset ...
func (r *RedisServer) Hset(key, field string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("hset", key, field, value))
	return
}

//Hexists ...
func (r *RedisServer) Hexists(key, field string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("hexists", key, field))
	return
}

//Hdel ...
func (r *RedisServer) Hdel(key, field string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("hdel", key, field))
	return
}

//Hkeys ...
func (r *RedisServer) Hkeys(key string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.String(conn.Do("hkeys", key))
	return
}

//Sadd ...
func (r *RedisServer) Sadd(key string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("sadd", key, value))
	return
}

//Smembers ...
func (r *RedisServer) Smembers(key string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Strings(conn.Do("smembers", key))
	return
}

//RemSet ...
func (r *RedisServer) RemSet(key string,value interface{}) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("srem", key, value))
	return
}

//Zrange ...
func (r *RedisServer) Zrange(key string, startIndex, endIndex int, isWithScores bool) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	if isWithScores {
		data, err = redis.Int64(conn.Do("zrange", key, startIndex, endIndex, "withscores"))
	} else {
		data, err = redis.Int64(conn.Do("zrange", key, startIndex, endIndex))
	}
	return
}

//Zadd ...
func (r *RedisServer) Zadd(key string,value interface{}, score int) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("zadd", key, score, value))
	return
}

func (r *RedisServer) TTL(key string) (data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("ttl", key))
	return
}

func (r *RedisServer) Exists(key string)(data interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	data, err = redis.Int64(conn.Do("exists", key))
	return
}