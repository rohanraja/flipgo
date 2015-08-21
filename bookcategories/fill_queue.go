package bookcategories

import "../crawlsave"
import "gopkg.in/mgo.v2/bson"
import "fmt"
import "encoding/json"

import "github.com/fatih/color"

func GetSavedParentObjects() (results []ParsedObject) {

	session, err := crawlsave.GetMongoConn()

	if err != nil {
		panic("Error in Mongo Connection")
	}

	c := session.DB(crawlsave.MONGO_DBNAME).C("categories")

	err = c.Find(bson.M{}).Sort("-_id").All(&results)

	return

}

func GetNextEnqueuingList(obj *ParsedObject) (outArr [][]string) {

	js, err := json.Marshal(obj)
	if err != nil {
		panic("Error in Marshalling")
	}

	jstr := string(js)
	step := 19

	for i := 1; i < 1521; {

		prUrl := getUrlFromSidNum(obj.Sid, i)

		prId := fmt.Sprintf("%s_%d", obj.Sid, i)

		var outStr []string
		outStr = append(outStr, jstr, prUrl, prId, obj.Url, obj.Sid)
		outArr = append(outArr, outStr)

		i = i + step
	}

	return
}

func GetFullEnqueueList(results []ParsedObject) (outArr [][]string) {

	for i := 0; i < len(results); i++ {

		newOutArr := GetNextEnqueuingList(&results[i])
		outArr = append(outArr, newOutArr...)
	}
	return
}

func EnqueueNextList(scrapeOut [][]string) {

	for i := 0; i < len(scrapeOut); i++ {

		newUrl := scrapeOut[i][1]
		newId := scrapeOut[i][2]

		crawlsave.Enqueue_NextCrawl(newUrl, newId, "resque:queue:bookcrawl", "Bookcrawl") // Todo : Catch Error
		c := color.New(color.FgGreen)
		c.Printf("\nEnqueued: %s", newId)
	}

}
