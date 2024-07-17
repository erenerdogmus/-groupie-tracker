// Package routes provides functions for handling HTTP routes and requests.
// routes/routes.go
package routes

import (
	"fmt"
	. "github.com/erenerdogmus/handlers"
	"net/http"
)

// HandleRequests handles all incoming HTTP requests and routes them to the appropriate controller functions.
func HandleRequests() {
	fmt.Println("Starting Server at Port 8045")
	fmt.Println("now open a browser and enter: localhost:8045 into the URL")
	http.HandleFunc("/", HandlerHomeFunc)
	http.HandleFunc("/artists", HandlerArtistFunc)
	http.HandleFunc("/search", HandlerSearchFunc)
	http.HandleFunc("/artistInfo", HandlerResaultFunc)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8045", nil)
}
