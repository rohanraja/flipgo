package main

import (
	"./crawlsave"
	"fmt"
)

func TryRedisConnection() bool {

	client := crawlsave.InitClient()
	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("REDIS ERROR: ", err)
		return false
	} else {
		fmt.Println("REDIS - OK")
		return true
	}

}

func testComponents() bool {

	redisTest := TryRedisConnection()

	return redisTest
}
