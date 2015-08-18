package main

import (
	"./bookcategories"
	// "./crawlsave"
	"fmt"
	"github.com/benmanns/goworker"
)

func main() {

	bookcategories.RegisterWorkers()
	// crawlsave.RegisterWorkers()

	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}

}
