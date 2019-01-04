package main

import (
"fmt"
"github.com/gomodule/redigo/redis"
)

func main() {
	options := redis.DialPassword("20081021Luck")
	c, err := redis.Dial("tcp", "123.59.136.131:6379", options)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("Connect success")
	result, err := redis.String(c.Do("GET", "test"))
	if err != nil {
		fmt.Println("redis get failed", err)
		return
	}
	fmt.Println(result)
}