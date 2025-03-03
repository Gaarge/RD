package main

import (
	"html/template"
	"log"
	"net/http"
)

type cssDate struct {
	CSS string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data cssDate) {
	templates, err := template.ParseFiles(
		"../templates/base.html",
		"../templates/header.html",
		"../templates/"+tmpl+".html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Fatal(err)
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register", cssDate{CSS: "login.css"})
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", cssDate{CSS: "about.css"})
}

func bad_pass(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "bad_pass", cssDate{CSS: "login.css"})
}

func login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", cssDate{CSS: "login.css"})
}

func shop(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "shop", cssDate{CSS: "shop.css"})
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", cssDate{CSS: "index.css"})
}
func lkp(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "lkp", cssDate{CSS: "lkp.css"})
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/lkp", lkp)
	http.HandleFunc("/about", about)
	http.HandleFunc("/bad_pass", bad_pass)

	staticDir := http.Dir("../static")
	staticHandler := http.FileServer(staticDir)
	staticHandler = http.StripPrefix("/static/", staticHandler)
	http.Handle("/static/", staticHandler)

	http.ListenAndServe(":8080", nil)
}
