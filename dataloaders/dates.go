package dataloaders

import (
	"encoding/json"
	"fmt"
	. "github.com/erenerdogmus/models"
	"net/http"
	"strings"
)

// Dates retrieves date data for a given artist ID from an external API and returns it as a DatesStruct.
func Dates(id string) DatesStruct {
	var DatesInfo = DatesStruct{}
	response, err := http.Get("https://github.com/erenerdogmuss.herokuapp.com/api/dates/" + id)
	if err != nil {
		fmt.Printf("HTTP request failed: %s", err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&DatesInfo)
	if err != nil {
		fmt.Printf("Failed to decode JSON response: %s", err)
	}
	RemoveStarFromDates(DatesInfo.Dates)
	return DatesInfo
}

// RemoveStarFromDates removes "*" character from each date string in a slice.
func RemoveStarFromDates(dates []string) {
	for i, date := range dates {
		if strings.Contains(date, "*") {
			dates[i] = strings.Replace(date, "*", "", -1)
		}
	}
}
