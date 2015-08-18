package book

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func Scrape(url2, pid string) (outArr [][]string, err error) {

	url_get := "http://flipkart.com" + url2
	doc, err := goquery.NewDocument(url_get)

	if err != nil {
		fmt.Println(err)
		return
	}
	prods := doc.Find("a.lu-title")

	prods.Each(func(i int, s *goquery.Selection) {

		url3 := strings.TrimSpace(s.AttrOr("href", "")) //.strip
		u, _ := url.Parse("http://flipkart.com" + url3)
		m, _ := url.ParseQuery(u.RawQuery)
		newpid := m["pid"][0]

		obj := ParsedObject{url3, newpid}
		js, err := json.Marshal(obj)
		if err != nil {
			return
		}

		jstr := string(js)
		var outStr []string

		outStr = append(outStr, jstr, url3, newpid)
		outArr = append(outArr, outStr)
	})

	return
}
