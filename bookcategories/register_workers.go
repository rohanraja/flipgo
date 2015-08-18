package bookcategories

import (
	"../crawlsave"
)

func RegisterWorkers() {

	crawlInfo := crawlsave.CrawlNode{}

	crawlInfo.COLL_NAME = "categories"
	crawlInfo.REDIS_SET_NAME = "saved_categories"
	crawlInfo.QUEUE_NAME = "resque:queue:catsavedb"
	crawlInfo.QUEUE_NEXT = "resque:queue:bookcrawl"
	crawlInfo.CLASS_NAME = "Catsavedb"
	crawlInfo.CLASS_NEXT = "Bookcrawl"
	crawlInfo.CLASS_PARSE = "Catcrawl"

	crawlInfo.QUEUE_PARSE = "resque:queue:catcrawl"
	crawlInfo.IS_RECURSIVE = true

	crawlInfo.Scrape = Scrape
	crawlInfo.GetMongoObj = getMongoObj

	crawlsave.RegisterWorkers(&crawlInfo)

}
