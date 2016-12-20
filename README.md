Redia Operation
===

基于redigo 封装读写操作


Installation
----

<pre><code>
go get github.com/johnhjwsosd/redis-operation/redisoper
</code></pre>



Example
----------
<pre><code>
    redis := redisoper.NewRedis("127.0.0.1:6379", "123")
	pool := redis.NewPool()
	res, err := redis.GetData(pool, "t2", "set")
    	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
</code></pre>