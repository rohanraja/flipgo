package bookcategories

import (
	// "encoding/json"
	"../crawlsave"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getUrlFromSidNum(sid string, num int) string {

	baseUrl := "/lc/pr/pv1/spotList1/spot1/productList"

	baseUrl = baseUrl + "?sid=" + sid + "&start=" + fmt.Sprintf("%d", num)

	return baseUrl

}

func Scrape(url, pid string) (outArr [][]string, err error) {

	url_get := "http://flipkart.com" + url

	doc, err := crawlsave.FetchHTMLDoc(url_get)

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

		step := 19

		for i := 1; i < 1521; {

			prUrl := getUrlFromSidNum(sid, i)
			prId := fmt.Sprintf("%s_%d", sid, i)

			var outStr []string
			outStr = append(outStr, jstr, prUrl, prId, url, sid)
			outArr = append(outArr, outStr)

			i = i + step
		}

	})

	return
}
