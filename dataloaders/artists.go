package dataloaders

import (
	"encoding/json"
	"fmt"
	. "github.com/erenerdogmus/models"
	"net/http"
)

// ArtistData retrieves artist data from an external API and returns it as a slice of Artist Structs.
func ArtistData() []ArtistStruct {
	var ArtistInfo = []ArtistStruct{}
	response, err := http.Get("https://github.com/erenerdogmuss.herokuapp.com/api/artists")
	if err != nil {
		fmt.Printf("HTTP request failed: %s", err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&ArtistInfo)
	if err != nil {
		fmt.Printf("Failed to decode JSON response: %s", err)
	}
	return ArtistInfo
}
