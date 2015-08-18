package bookinfo

import (
	"../crawlsave"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func bookInfoSaveWorker(queue string, args ...interface{}) error {
	intf := args[0]

	bijson, ok := intf.(string)
	err := errors.New("ERROR!")

	if ok == false {
		return err
	}

	var bi Bookinfo

	err = json.Unmarshal([]byte(bijson), &bi)

	if err != nil {
		fmt.Println(err)
		return err
	}

	pid := bi.Pid

	alreadyDone := crawlsave.CheckIfAlreadyDone(pid, REDIS_SET_NAME) // Todo : Catch Error

	if alreadyDone == true {
		fmt.Println("Skipping")
		return nil
	}
	st := time.Now()
	err = crawlsave.SaveToDB(&bi, COLL_NAME)
	elapsed := time.Since(st)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("dbsave took %s\n", elapsed)

	// fmt.Println(bi)
	crawlsave.MarkAsDone(pid, REDIS_SET_NAME)

	return nil
}
