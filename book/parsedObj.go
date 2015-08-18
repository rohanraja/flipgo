package book

import (
	"encoding/json"
)

type ParsedObject struct {
	Url string `json:"url"`
	Pid string `json:"pid"`
}

func getMongoObj(jsString string) (interface{}, error) {

	var bi ParsedObject

	err := json.Unmarshal([]byte(jsString), &bi)

	return &bi, err

}
