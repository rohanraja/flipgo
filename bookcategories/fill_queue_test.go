package bookcategories

import (
	// "encoding/json"
	// "github.com/benmanns/goworker"
	"testing"
)

import "github.com/fatih/color"

func TesTRetrieveingMongoDocs(t *testing.T) {

	c := color.New(color.FgGreen)
	results := GetSavedParentObjects()
	c.Println(results[0].Sid)

}

func TesTGeneratingList(t *testing.T) {

	results := GetSavedParentObjects()
	outArr := GetNextEnqueuingList(&results[0])

	c := color.New(color.FgGreen)
	c.Println(outArr[0])

}

func TesTEnqueueingList(t *testing.T) {

	results := GetSavedParentObjects()
	outArr := GetNextEnqueuingList(&results[0])

	EnqueueNextList(outArr)

}

func TesTEnqueueingFullList(t *testing.T) {

	results := GetSavedParentObjects()
	outArr := GetFullEnqueueList(results)

	EnqueueNextList(outArr)

}
