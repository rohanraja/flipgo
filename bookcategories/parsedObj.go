package bookcategories

import (
	"encoding/json"
)

type ParsedObject struct {
	Url         string `json:"url"`
	Sid         string `json:"sid"`
	Numproducts string `json:"numproducts"`
	Title       string `json:"title"`
}

func getMongoObj(jsString string) (interface{}, error) {

	var bi ParsedObject

	err := json.Unmarshal([]byte(jsString), &bi)

	return &bi, err

}
