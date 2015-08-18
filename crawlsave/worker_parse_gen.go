package crawlsave

import (
	"errors"
	"fmt"
)

func GetWorkerFunc(redis_set_name,
	queue_db,
	class_db,
	queue_crawl,
	class_crawl string,
	scrape func(string, string) ([][]string, error)) (f func(string, ...interface{}) error) {
	// extractNextCrawlParams func(string) (string, string)) (f func(string, ...interface{}) error) {

	f = func(queue string, args ...interface{}) error {
		url, ok := args[0].(string)
		id, ok := args[1].(string)

		fmt.Println(url)

		if ok == false {
			return errors.New("Cannot parse input arguments")
		}

		alreadyDone := CheckIfAlreadyDone(id, redis_set_name) // tODo : Catch Error

		if alreadyDone == true {
			fmt.Println("Skipping")
			return nil
		}

		scrapeOut, err := scrape(url, id)

		if err != nil {
			fmt.Println("Error in parsing!")
			return err
		}

		for i := 0; i < len(scrapeOut); i++ {

			jstr := scrapeOut[i][0]
			newUrl := scrapeOut[i][1]
			newId := scrapeOut[i][2]

			Enqueue_NextCrawl(jstr, id, queue_db, class_db) // Todo : Catch Error

			Enqueue_NextCrawl(newUrl, newId, queue_crawl, class_crawl) // Todo : Catch Error

		}

		fmt.Println("ParsingDone!")

		return nil
	}

	return
}
