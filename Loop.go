package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"time"



)

func main(){

	for i :=0; i< 100; i++{

		if !(i >= 100) {
			fmt.Println("Counter is i", i)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("You have reached i", i)


	}

	_ = connectToRedis()
}

func connectToRedis() (conn *redis.Client) {
	conn = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
	return
}

func SaveToRedis(Key string, data interface{}, client *redis.Client) (err error) {
	expiryTime := time.Duration(time.Minute * 2)

	jsonData, _ := json.Marshal(data)

	if err = client.Set(Key, jsonData, expiryTime).Err(); err != nil {
		log.Println("Unable to save data to redis because ", err)
		return
	}
	return
}

func ReadFromRedis(Key string, client *redis.Client) (data string, err error) {
	data, err = client.Get(Key).Result()
	if err != nil {
		fmt.Println("Unable to read ", Key, " because ", err)
		return
	}
	return
}

