package bookcategories

import (
	// "encoding/json"
	"fmt"
	// "github.com/benmanns/goworker"
	"testing"
)

func TestParsingAUrl(t *testing.T) {

	url := "/books/pr?sid=bks"

	out, err := Scrape(url, "123")

	fmt.Println(out, err)

}