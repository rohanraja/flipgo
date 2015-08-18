package crawlsave

import (
	"github.com/kavu/go-resque"         // Import this package
	_ "github.com/kavu/go-resque/godis" // Use godis driver
	"github.com/simonz05/godis/redis"   // Redis client from godis package
)

func Enqueue_NextCrawl(url, id, queue_name, class_name string) {
	var err error

	if queue_name == "none" {
		return
	}

	client := redis.New(REDIS_ENQUEUE_HOST, 0, "")       // Create new Redis client to use for enqueuing
	enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance

	// Enqueue the job into the "go" queue with appropriate client
	_, err = enqueuer.Enqueue(queue_name, class_name, url, id)
	if err != nil {
		panic(err)
	}
}
