package crawlsave

import (
	"gopkg.in/redis.v3"
)

func InitClient() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return
}

func CheckIfAlreadyDone(pid, setname string) (out bool) {

	client := InitClient()
	isMem := client.SIsMember(setname, pid)
	out = isMem.Val()

	return
}

func MarkAsDone(pid, setname string) {

	client := InitClient()
	client.SAdd(setname, pid)
}
