package dataloaders

import (
	"encoding/json"
	"fmt"
	. "github.com/erenerdogmus/models"
	"net/http"
)

// relation retrieves relation data for a given artist ID from an external API and returns it as a Relation Struct.
func Relation(id string) RelationStruct {
	var RelationInfo = RelationStruct{}
	response, err := http.Get("https://github.com/erenerdogmuss.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Printf("HTTP request failed: %s", err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&RelationInfo)
	if err != nil {
		fmt.Printf("Failed to decode JSON response: %s", err)
	}
	return RelationInfo
}
