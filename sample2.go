//go:build ignore

package main

import (
	"html/template"
	"net/http"
)

type Site struct {
	Title       string
	Description string
}

var site = &Site{
	Title:       "My WebSite",
	Description: "My WebSite Description",
}

func main() {
	http.HandleFunc("/", pageHandler) // http://localhost:3070/

	http.ListenAndServe(":3070", nil)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	tmpl.ExecuteTemplate(w, "layout", site)
}
