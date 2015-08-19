package crawlsave

// var ITEM_LIST_DB string = "saved_bookinfos"

import (
	"fmt"
	"os"
)

var MainserverIP string = "10.109.11.216"

var REDIS_HOST string = MainserverIP + ":6379"
var MONGO_HOST string = "localhost"
var MONGO_DBNAME string = "flipgo_dev"
var REDIS_ENQUEUE_HOST string = "tcp:" + MainserverIP + ":6379"

func init() {

	if len(os.Args) > 1 {
		fmt.Println("Setting custom redis server")
		serverIp := os.Args[1]
		MainserverIP = serverIp
		REDIS_HOST = MainserverIP + ":6379"
		REDIS_ENQUEUE_HOST = "tcp:" + MainserverIP + ":6379"
	}

}
