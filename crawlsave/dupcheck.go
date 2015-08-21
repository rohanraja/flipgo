package crawlsave

import (
	"errors"
	"gopkg.in/redis.v3"
	"time"
)

import "github.com/fatih/color"

var client *redis.Client

func InitClient() (clientp *redis.Client) {
	clientp = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return
}

func CheckIfAlreadyDone(pid, setname string) (out bool) {

	retries := 15
	err := errors.New("Init Error")

	for retries > 0 && err != nil {
		out, err = client.SIsMember(setname, pid).Result()
		if err != nil {
			c := color.New(color.FgYellow)
			c.Println(retries, "\nREDIS! Error in queueing on redis, Retrying..", err)
			time.Sleep(10 * time.Second)
			client = InitClient()
		}

		retries = retries - 1
	}

	if err != nil {

		c := color.New(color.FgYellow)
		c.Println("ENQUEING ERROR: ", err)
		panic(err)
	}

	return
}

func MarkAsDone(pid, setname string) {

	retries := 15
	err := errors.New("Init Error")

	for retries > 0 && err != nil {
		_, err = client.SAdd(setname, pid).Result()
		if err != nil {
			c := color.New(color.FgRed)
			c.Println(retries, "\nREDIS! Error in queueing on redis, Retrying..", err)
			time.Sleep(10 * time.Second)
			client = InitClient()
		}

		retries = retries - 1
	}

	if err != nil {

		c := color.New(color.FgRed)
		c.Println("ENQUEING ERROR: ", err)
		panic(err)
	}

	return
}

func init() {

	client = InitClient()

}
