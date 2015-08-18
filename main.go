package main

import (
	"./bookinfo"
	// "./crawlsave"
	"fmt"
	"github.com/benmanns/goworker"
)

func main() {

	bookinfo.RegisterWorkers()
	// crawlsave.RegisterWorkers()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}

}
