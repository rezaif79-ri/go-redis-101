package model

import "github.com/gomodule/redigo/redis"

type User struct {
	Key  string
	Data UserData
}

type UserData struct {
	UserName    string
	PhoneNumber string
	Height      float64
	Cash        int64
}

func RedisGetHGetUserData(redisConn redis.Conn, key string) (User, error) {
	rep, err := redis.StringMap(redisConn.Do("HGETALL", key))
	if err != nil {
		return User{}, err
	}
	userName, _ := redis.String(rep["user_name"], nil)
	phoneNumber, _ := redis.String(rep["phone_number"], nil)
	height, _ := redis.Float64(rep["height"], nil)
	cash, _ := redis.Int64(rep["cash"], nil)

	return User{
		Key: key,
		Data: UserData{
			UserName:    userName,
			PhoneNumber: phoneNumber,
			Height:      height,
			Cash:        cash,
		},
	}, err
}

func RedisHSetUserData(redisConn redis.Conn, key string, data UserData) (interface{}, error) {
	rep, err := redisConn.Do("HSET", key,
		"user_name", data.UserName,
		"phone_number", data.PhoneNumber,
		"height", data.Height,
		"cash", data.Cash,
	)
	return rep, err
}
