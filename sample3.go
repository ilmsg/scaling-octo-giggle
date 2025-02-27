//go:build ignore

package main

import (
	"html/template"
	"net/http"
	"path/filepath"
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
	Render(w, site, []string{"index.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
