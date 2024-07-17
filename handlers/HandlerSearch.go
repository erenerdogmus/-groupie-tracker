package handlers

import (
	handleerror "github.com/erenerdogmus/handleError"
	"github.com/erenerdogmus/utils"
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("views/searchbar.html"))

func HandlerSearchFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		handleerror.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	queryParams := r.URL.Query()
	hasQuery := len(queryParams) > 0
	searchTerm := r.FormValue("search")
	data := utils.CollectData(strings.ToUpper(searchTerm))
	if hasQuery {
		data = utils.FilterControl(queryParams)
	}
	if err := tmpl.Execute(w, data); err != nil {
		handleerror.ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	queryParams := make(map[string][]string)
// 	for key, values := range r.URL.Query() {
// 		queryParams[key] = values
// 	}
// 	filteredProducts := filterProducts(queryParams)
// 	for _, p := range filteredProducts {
// 		fmt.Fprintf(w, "%+v\n", p)
// 	}
// }
