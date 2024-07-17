package dataloaders

import (
	"encoding/json"
	"fmt"
	. "github.com/erenerdogmus/models"
	"net/http"
)

// Location retrieves location data for a given artist ID from an external API and returns it as a LocationStruct.
func Location(id string) LocationStruct {
	var locationInfo LocationStruct
	response, err := http.Get("https://github.com/erenerdogmuss.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Printf("HTTP request failed: %s\n", err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&locationInfo)
	if err != nil {
		fmt.Printf("Failed to decode JSON response: %s", err)
	}
	return locationInfo
}
