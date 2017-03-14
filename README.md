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

	res, err = redis.Set("t1", "22")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err := redis.Get("t1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res.([]byte)))
	}

</code></pre>