package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

var Client redis.Conn = nil

func InitRedis() {
	options := redis.DialPassword("20081021Luck")
	client, err := redis.Dial("tcp", "123.59.136.131:6379", options)
	if err != nil {
		log.Println("connect redis fail:", err)
		return
	}
	Client = client
	log.Println("Connect redis success")
}