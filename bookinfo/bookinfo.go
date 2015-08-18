package bookinfo

import (
	"encoding/json"
)

type Bookinfo struct {
	Title       string            `json:"title"`
	Price       string            `json:"price"`
	Metadata    map[string]string `json:"metadata"`
	Imgs        []string          `json:"imgs"`
	Mainimg     string            `json:"mainimg"`
	Url         string            `json:"url"`
	Pid         string            `json:"pid"`
	Author      string            `json:"author"`
	Categories  string            `json:"categories"`
	Description string            `json:"description"`
}

func getMongoObj(jsString string) (interface{}, error) {

	var bi Bookinfo

	err := json.Unmarshal([]byte(jsString), &bi)

	return &bi, err

}
