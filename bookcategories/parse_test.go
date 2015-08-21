package bookcategories

import (
	// "encoding/json"
	"fmt"
	// "github.com/benmanns/goworker"
	"testing"
)

func TesTParsingAUrl(t *testing.T) {

	url := "/books/pr?sid=bks"

	out, err := Scrape(url, "123")

	fmt.Println(out, err)

}

func TesTUrlGeneration(t *testing.T) {

	a := getUrlFromSidNum("bks", 93)
	fmt.Println(a)

}
