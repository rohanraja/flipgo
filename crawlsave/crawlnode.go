package crawlsave

type CrawlNode struct {
	COLL_NAME      string
	REDIS_SET_NAME string
	QUEUE_NAME     string
	QUEUE_NEXT     string
	CLASS_NAME     string
	CLASS_PARSE    string
	CLASS_NEXT     string
	Scrape         func(string, string) ([][]string, error)
	GetMongoObj    func(string) (interface{}, error)

	IS_RECURSIVE bool
	QUEUE_PARSE  string
}
