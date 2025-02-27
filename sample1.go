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
	http.HandleFunc("/", pageHandler)       // http://localhost:3070/
	http.HandleFunc("/hello", helloHandler) // http://localhost:3070/hello

	http.ListenAndServe(":3070", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello."))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/page1.html")
	tmpl.Execute(w, site)
}
