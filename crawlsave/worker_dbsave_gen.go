package crawlsave

import (
	"errors"
	"fmt"
)

func GetWorkerDBSaveFunc(crawlinfo *CrawlNode) (f func(string, ...interface{}) error) {

	f = func(queue string, args ...interface{}) error {

		bijson, ok := args[0].(string)
		id, ok := args[1].(string)

		err := errors.New("ERROR!")

		if ok == false {
			fmt.Println("Error in passing arguments")
			return err
		}

		alreadyDone := CheckIfAlreadyDone(id, crawlinfo.REDIS_SET_NAME) // tODo : Catch Error

		if alreadyDone == true {
			fmt.Println("Skipping")
			return nil
		}

		saveBuffer, err := crawlinfo.GetMongoObj(bijson)

		if err != nil {
			fmt.Println("Error in Marshalling!")
			return err
		}

		err = SaveToDB(saveBuffer, crawlinfo.COLL_NAME)

		if err != nil {
			fmt.Println("Error in saving!")
			return err
		}

		MarkAsDone(id, crawlinfo.REDIS_SET_NAME)
		fmt.Println("SaveDBDone!")

		return nil
	}

	return
}
