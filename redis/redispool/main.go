package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//初始化函数  自己就会调用的   而且无需大写
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,   //0代表是不受限制的
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")

		},
	}

}

func main() {
	//从redis pool中获取一个conn
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set is failed err:", err)

	}
	recv, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("recv is failed err:", err)
	}
	fmt.Println(recv)
}
