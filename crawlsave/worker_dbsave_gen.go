package crawlsave

import (
	"errors"
	"fmt"
)

func GetWorkerDBSaveFunc(redis_set_name,
	coll_name string,
	getMongoObject func(string) (interface{}, error)) (f func(string, ...interface{}) error) {

	f = func(queue string, args ...interface{}) error {

		bijson, ok := args[0].(string)
		id, ok := args[1].(string)

		err := errors.New("ERROR!")

		if ok == false {
			fmt.Println("Error in passing arguments")
			return err
		}

		alreadyDone := CheckIfAlreadyDone(id, redis_set_name) // tODo : Catch Error

		if alreadyDone == true {
			fmt.Println("Skipping")
			return nil
		}

		saveBuffer, err := getMongoObject(bijson)

		if err != nil {
			fmt.Println("Error in Marshalling!")
			return err
		}

		err = SaveToDB(saveBuffer, coll_name)

		if err != nil {
			fmt.Println("Error in saving!")
			return err
		}

		MarkAsDone(id, redis_set_name)
		fmt.Println("SaveDBDone!")

		return nil
	}

	return
}
