package crawlsave

import (
	// "encoding/json"
	"fmt"
	"testing"
)

func TestFetchingaUrl(t *testing.T) {

	url := "htp://www.f1lipkart.com/books/educational-and-professional/academic-texts/law/administrative-law-regulatory-practice/pr?p%5B%5D=facets.price_range%255B%255D%3DRs.%2B1001%2Band%2BAbove&p%5B%5D=sort%3Dprice_asc&sid=bks%2Cenp%2Cq2s%2C71w%2Citf&pincode=721302&filterNone=true"
	fmt.Println("Testing fetching URL: ", url)

	doc, _ := FetchHTMLDoc(url)

	fmt.Println(doc.Find("title").Text())

}
