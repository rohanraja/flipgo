package main

import (
	"./book"
	"./bookcategories"
	"./bookinfo"
	// "./crawlsave"
	"fmt"
	"github.com/benmanns/goworker"
)

func main() {

	bookcategories.RegisterWorkers()
	book.RegisterWorkers()
	bookinfo.RegisterWorkers()
	// crawlsave.RegisterWorkers()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}

}
