package crawlsave

import (
	// "encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

import "github.com/fatih/color"

func FetchHTMLDoc(url string) (doc *goquery.Document, err error) {

	retries := 15

	err = errors.New("Init Errir")
	for retries > 0 && err != nil {

		doc, err = goquery.NewDocument(url)
		if err != nil {
			c := color.New(color.FgMagenta)
			c.Println(retries, "RETRYING! Error in fetching http:", err)
			time.Sleep(10 * time.Second)
		}
		retries = retries - 1

	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if retries < 14 {
		c := color.New(color.FgGreen)
		c.Printf("\nRecovered from http error in %d tries\n", 14-retries)
	}
	return
}
