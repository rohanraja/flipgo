package crawlsave

import (
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "log"
)

func GetMongoConn() (session *mgo.Session, err error) {

	session, err = mgo.Dial(MONGO_HOST)
	return
}

func SaveToDB(bi interface{}, collname string) error {

	session, err := GetMongoConn()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(MONGO_DBNAME).C(collname)

	err = c.Insert(bi)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
