package crawlsave

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kavu/go-resque"         // Import this package
	_ "github.com/kavu/go-resque/godis" // Use godis driver
	"github.com/simonz05/godis/redis"   // Redis client from godis package
	"time"
)
import "github.com/fatih/color"

type jobArg interface{}

type job struct {
	Class string   `json:"class"`
	Args  []jobArg `json:"args"`
}

func getQueueString(jobClass string, args ...jobArg) string {

	jobJSON, err := json.Marshal(&job{jobClass, args})
	if err != nil {
		panic(err)
	}

	return string(jobJSON)

}

func Enqueue_NextCrawl(url, id, queue_name, class_name string) {

	if queue_name == "none" {
		return
	}
	queue_obj_str := getQueueString(class_name, url, id)
	// queue_obj_str := fmt.Sprintf("{\"class\":\"%s\",\"args\":[\"%s\",\"%s\"]}", class_name, url, id)

	var err error

	retries := 15
	err = errors.New("Init Error")

	for retries > 0 && err != nil {
		_, err = client.RPush(queue_name, queue_obj_str).Result()

		if err != nil {
			c := color.New(color.FgYellow)
			c.Println(retries, "\nREDIS! Error in queueing on redis, Retrying..", err)
			time.Sleep(10 * time.Second)
			client = InitClient()
		}

		retries = retries - 1
	}

	if err != nil {
		fmt.Println("ENQUEING ERROR: ", err)
		panic(err)
	}

	if retries < 14 {
		c := color.New(color.FgCyan)
		c.Printf("\nRecovered from redis error in %d tries\n", 14-retries)
	}

	return
}

func old_Enqueue_NextCrawl(url, id, queue_name, class_name string) {

	if queue_name == "none" {
		return
	}

	var err error

	retries := 15
	err = errors.New("Init Errir")
	for retries > 0 && err != nil {

		client := redis.New(REDIS_ENQUEUE_HOST, 0, "")       // Create new Redis client to use for enqueuing
		enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance

		_, err = enqueuer.Enqueue(queue_name, class_name, url, id)

		if err != nil {
			c := color.New(color.FgYellow)
			c.Println(retries, "REDIS! Error in queueing on redis", err)
			time.Sleep(10 * time.Second)
		}

		retries = retries - 1
	}

	if err != nil {
		fmt.Println("ENQUEING ERROR: ", err)
		panic(err)
	}

	if retries < 14 {
		c := color.New(color.FgCyan)
		c.Printf("\nRecovered from redis error in %d tries\n", 14-retries)
	}

	return

}
