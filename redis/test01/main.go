package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connnect failed err:", err)
		return
	}
	//一定要记住关闭
	
	defer conn.Close()
	//设置存档的值
	_, err = conn.Do("set", "name", "Tom:自信即巅峰")
	if err != nil {
		fmt.Println("set is failed err:", err)
	}
	//获取存入的值
	rec, err := conn.Do("get", "name")
	if err != nil {
		fmt.Println("get is failed err")

	}
	rec, err = redis.String(rec, err)
	if err != nil {
		fmt.Println("redis.string is failed err:", err)
	}

	fmt.Println(rec) //ok
}
