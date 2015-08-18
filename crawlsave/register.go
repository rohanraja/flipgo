package crawlsave

import (
	"github.com/benmanns/goworker"
)

func RegisterWorkers(crawlinfo *CrawlNode) {

	worker_parse := GetWorkerFunc(
		crawlinfo.REDIS_SET_NAME,
		crawlinfo.QUEUE_NAME,
		crawlinfo.CLASS_NAME,
		crawlinfo.QUEUE_NEXT,
		crawlinfo.CLASS_NEXT,
		crawlinfo.Scrape,
	)

	worker_save := GetWorkerDBSaveFunc(
		crawlinfo.REDIS_SET_NAME,
		crawlinfo.COLL_NAME,
		crawlinfo.GetMongoObj,
	)

	goworker.Register(crawlinfo.CLASS_PARSE, worker_parse)
	goworker.Register(crawlinfo.CLASS_NAME, worker_save)
}
