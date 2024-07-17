// Package handleerror provides functions for handling and rendering error messages and pages.
// handleError/handlerError.go
package handleerror

import (
	"fmt"
	. "github.com/erenerdogmus/models"
	"net/http"
	"text/template"
)

// ErrorHandler handles error messages and renders corresponding error pages.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("views/errorPage.html")
	if err != nil {
		fmt.Fprintf(w, "HTTP status %d: Internal Server Error - missing errorPage.html file", http.StatusInternalServerError)
		return
	}
	var em string
	switch status {
	case http.StatusNotFound:
		em = "HTTP status 404: Page Not Found"
	case http.StatusInternalServerError:
		em = "HTTP status 500: Internal Server Error"
	case http.StatusBadRequest:
		em = "HTTP status 400: Bad Request! Please select artist from the Home Page"
	default:
		em = fmt.Sprintf("HTTP status %d: Unknown Error", status)
	}
	p := Error{ErrorNum: status, ErrorMes: em}
	t.Execute(w, p)
}
