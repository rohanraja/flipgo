package book

import (
	// "encoding/json"
	"fmt"
	// "github.com/benmanns/goworker"
	"testing"
)

func TestParsingAUrl(t *testing.T) {

	url := "/lc/pr/pv1/spotList1/spot1/productList?sid=bks&pincode=721302&filterNone=true&start=41&ajax=true&_=1439910087539"

	out, err := Scrape(url, "123")

	fmt.Println(out, err)

}
