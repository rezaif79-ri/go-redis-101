package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/rezaif79-ri/go-redis-101/model"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Panic(err)
	}

	rep, err := model.RedisHSetUserData(conn, "user:2", model.UserData{
		UserName:    "Abdul",
		PhoneNumber: "0822371238",
		Height:      199.8,
		Cash:        300000,
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(rep)

	userData, err := model.RedisGetHGetUserData(conn, "user:2")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(userData)
}
