package bookinfo

import (
	"../crawlsave"
	"errors"
	"fmt"
	"time"
)

func bookInfoWorker(queue string, args ...interface{}) error {

	intf := args[0]
	mapp := intf.(map[string]interface{})
	url, ok := mapp["url"].(string)
	pid, ok := mapp["pid"].(string)

	err := errors.New("ERROR!")

	if ok == false {
		return err
	}

	alreadyDone := crawlsave.CheckIfAlreadyDone(pid, REDIS_SET_NAME) // Todo : Catch Error

	if alreadyDone == true {
		fmt.Println("Skipping")
		return nil
	}

	st := time.Now()
	jstr := Scrape(url, pid) // Todo : Catch Error
	elapsed := time.Since(st)

	fmt.Printf("network took %s\n", elapsed)

	crawlsave.Enqueue_DBSave(jstr, QUEUE_NAME, CLASS_NAME) // Todo : Catch Error
	return nil
}
