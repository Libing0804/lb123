package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dial is failed err:", err)
		return
	}
	defer conn.Close()
	//一次存放一个键值对
	// _, err = conn.Do("hset", "hero01", "name", "tom")
	// if err != nil {
	// 	fmt.Println("hset is failed err: ", err)
	// }
	// _, err = conn.Do("hset", "hero01", "age", 24)
	// if err != nil {
	// 	fmt.Println("hset is failed err: ", err)
	// }
	// _, err = conn.Do("hset", "hero01", "golang", "自信即巅峰")
	// if err != nil {
	// 	fmt.Println("hset is failed err: ", err)
	// }
	// rec1, err := redis.String(conn.Do("hget", "hero01", "name"))
	// if err != nil {
	// 	fmt.Println("failed err:", err)
	// }
	// fmt.Println(rec1)
	// rec2, err := redis.Int(conn.Do("hget", "hero01", "age"))
	// if err != nil {
	// 	fmt.Println("failed err:", err)
	// }
	// fmt.Println(rec2)
	// rec3, err := redis.String(conn.Do("hget", "hero01", "golang"))
	// if err != nil {
	// 	fmt.Println("failed err:", err)
	// }
	//一次存放多个hash的方式
	_, err = conn.Do("hmset", "hero01", "name", "jerry", "age", 25, "golang", "无畏即攀登")
	if err != nil {
		fmt.Println("hmset failed err:", err)
	}
	recv, err := redis.Strings(conn.Do("hmget", "hero01", "name", "age", "golang"))
	if err != nil {
		fmt.Println("hmget is failed err:", err)
	}
	for i, v := range recv {
		fmt.Printf("recv[%d]=%v\n", i, v)
	}
}
