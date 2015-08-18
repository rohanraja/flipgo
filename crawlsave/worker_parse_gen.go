package crawlsave

import (
	"errors"
	"fmt"
)

func GetWorkerFunc(crawlinfo *CrawlNode) (f func(string, ...interface{}) error) {
	// extractNextCrawlParams func(string) (string, string)) (f func(string, ...interface{}) error) {

	f = func(queue string, args ...interface{}) error {
		url, ok := args[0].(string)
		id, ok := args[1].(string)

		fmt.Println(url)

		if ok == false {
			return errors.New("Cannot parse input arguments")
		}

		// alreadyDone := CheckIfAlreadyDone(id, crawlinfo.REDIS_SET_NAME) // tODo : Catch Error
		//
		// if alreadyDone == true {
		// 	fmt.Println("Skipping")
		// 	return nil
		// }

		scrapeOut, err := crawlinfo.Scrape(url, id)

		if err != nil {
			fmt.Println("Error in parsing!")
			return err
		}

		for i := 0; i < len(scrapeOut); i++ {

			jstr := scrapeOut[i][0]
			newUrl := scrapeOut[i][1]
			newId := scrapeOut[i][2]

			Enqueue_NextCrawl(jstr, newId, crawlinfo.QUEUE_NAME, crawlinfo.CLASS_NAME) // Todo : Catch Error

			Enqueue_NextCrawl(newUrl, newId, crawlinfo.QUEUE_NEXT, crawlinfo.CLASS_NEXT) // Todo : Catch Error

		}

		fmt.Println("ParsingDone!")

		return nil
	}

	return
}
