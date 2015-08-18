package bookinfo

import (
	"../crawlsave"
)

func RegisterWorkers() {

	crawlInfo := crawlsave.CrawlNode{}

	crawlInfo.COLL_NAME = "bookinfos"
	crawlInfo.REDIS_SET_NAME = "saved_bookinfos"
	crawlInfo.QUEUE_NAME = "resque:queue:bookinfosavedb"
	crawlInfo.QUEUE_NEXT = "none"
	crawlInfo.CLASS_NAME = "Bookinfosavedb"
	crawlInfo.CLASS_NEXT = "none"
	crawlInfo.CLASS_PARSE = "Bookinfocrawl"

	crawlInfo.Scrape = Scrape
	crawlInfo.GetMongoObj = getMongoObj

	crawlsave.RegisterWorkers(&crawlInfo)

}
