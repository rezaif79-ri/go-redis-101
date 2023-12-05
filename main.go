package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Panic(err)
	}
	rep, err := conn.Do("HSET", "user:1", "user_name", "Reza IF", "phone_number", "0812371238", "height", 165.8, "cash", 10000)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(rep)
}
