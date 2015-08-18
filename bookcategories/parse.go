package bookcategories

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Scrape(url, pid string) (outArr [][]string, err error) {

	url_get := "http://flipkart.com" + url
	doc, err := goquery.NewDocument(url_get)

	if err != nil {
		fmt.Println(err)
		return
	}

	substores := doc.Find("#substores")

	allLi := substores.Find("li[class='store']")

	allLi.Each(func(i int, s *goquery.Selection) {
		url := strings.TrimSpace(s.Find("a").Eq(0).AttrOr("href", ""))
		sid := strings.TrimSpace(s.Find("a").Eq(0).AttrOr("sid", ""))
		num := strings.TrimSpace(s.Find("span").Eq(0).Text())
		title := strings.TrimSpace(s.Find("a").Text())

		obj := ParsedObject{url, sid, num, title}
		js, err := json.Marshal(obj)
		if err != nil {
			return
		}

		jstr := string(js)
		var outStr []string

		outStr = append(outStr, jstr, url, sid)
		outArr = append(outArr, outStr)

	})

	return
}
