package main

import (
	"./book"
	"./bookcategories"
	"./bookinfo"
	"./goworker"
	"fmt"
)

func main() {

	isWorking := testComponents()

	if isWorking == false {
		fmt.Println("Quitting!")
		return
	}

	bookcategories.RegisterWorkers()
	book.RegisterWorkers()
	bookinfo.RegisterWorkers()
	// crawlsave.RegisterWorkers()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}

}
