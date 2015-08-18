package crawlsave

import (
	"github.com/benmanns/goworker"
)

func RegisterWorkers(crawlinfo *CrawlNode) {

	worker_parse := GetWorkerFunc(crawlinfo)
	worker_save := GetWorkerDBSaveFunc(crawlinfo)

	goworker.Register(crawlinfo.CLASS_PARSE, worker_parse)
	goworker.Register(crawlinfo.CLASS_NAME, worker_save)
}
