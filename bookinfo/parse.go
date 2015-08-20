package bookinfo

import (
	"../crawlsave"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Scrape(url, pid string) (outArr [][]string, err error) {

	url_get := "http://flipkart.com" + url
	doc, err := crawlsave.FetchHTMLDoc(url_get)

	if err != nil {
		fmt.Println(err)
		return
	}

	title := doc.Find("h1.title").Text()
	price := doc.Find("span.selling-price").Text()
	mainimg := doc.Find(".mainImage img").Eq(0).AttrOr("data-src", "")
	author := doc.Find(".bookDetails a").Text()
	description := doc.Find(".description-text").Text()

	categories := strings.TrimSpace(doc.Find(".clp-breadcrumb").Text())
	categories = strings.Replace(categories, "\t", "", -1)
	categories = strings.Replace(categories, "\n", "", -1)
	categories = strings.Replace(categories, "\u003e", ",", -1)
	categories = strings.Replace(categories, "\u0026", ",", -1)

	metadata := map[string]string{}
	doc.Find(".specTable tr").Each(func(i int, s *goquery.Selection) {
		key := s.Find("td").Eq(0).Text()
		value := s.Find("td").Eq(1).Text()
		metadata[strings.TrimSpace(key)] = strings.TrimSpace(value)
	})

	imgs := []string{}
	doc.Find(".carouselContainer ul li").Each(func(i int, s *goquery.Selection) {

		im := s.Find(".thumb").Eq(0).AttrOr("style", "")
		imgs = append(imgs, im)

	})

	bi := Bookinfo{title, price, metadata, imgs, mainimg, url, pid, author, categories, description}

	js, err := json.Marshal(bi)
	if err != nil {
		return
	}

	var outStr []string

	outStr = append(outStr, string(js))
	outStr = append(outStr, "")
	outStr = append(outStr, pid)
	outArr = append(outArr, outStr)

	return
}
