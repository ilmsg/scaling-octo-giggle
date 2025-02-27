package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/public/", http.StripPrefix("/public/", fs))

	r.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/about", aboutHandler).Methods(http.MethodGet)

	r.HandleFunc("/auth/register", authRegisterHandler).Methods(http.MethodGet)
	r.HandleFunc("/auth/login", authLoginHandler).Methods(http.MethodGet)
	r.HandleFunc("/auth/profile", authProfileHandler).Methods(http.MethodGet)

	http.ListenAndServe(":3070", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"index.html"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"about.html"})
}

func authRegisterHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"auth/register.html"})
}

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"auth/login.html"})
}

func authProfileHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, site, []string{"auth/profile.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
