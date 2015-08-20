package crawlsave

// var ITEM_LIST_DB string = "saved_bookinfos"

import (
	"flag"
	// "fmt"
)

var MainserverIP string = "127.0.0.1"

var REDIS_HOST string = MainserverIP + ":6379"
var MONGO_HOST string = "localhost"
var MONGO_DBNAME string = "flipgo_21_8"
var REDIS_ENQUEUE_HOST string = "tcp:" + MainserverIP + ":6379"

func init() {

	rserver := flag.String("redis", "127.0.0.1", "Redis Server IP")
	flag.Parse()

	// if len(os.Args) > 1 {
	// fmt.Println("Setting custom redis server")
	// serverIp := os.Args[1]
	serverIp := *rserver
	MainserverIP = serverIp
	REDIS_HOST = MainserverIP + ":6379"
	REDIS_ENQUEUE_HOST = "tcp:" + MainserverIP + ":6379"
	// }

}
