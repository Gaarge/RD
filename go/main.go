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
		"../templates/footer.html",
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

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", cssDate{CSS: "index.css"})
}

func main() {

	http.HandleFunc("/", index)

	staticDir := http.Dir("../static")
	staticHandler := http.FileServer(staticDir)
	staticHandler = http.StripPrefix("/static/", staticHandler)
	http.Handle("/static/", staticHandler)

	/*

		mainDir := http.Dir("../templates")
		templatesHandler := http.FileServer(mainDir)
		http.Handle("/", templatesHandler)

		staticDir := http.Dir("../static")
		staticHandler := http.FileServer(staticDir)
		templatesHandler = http.StripPrefix("/static/", staticHandler)
		http.Handle("/static/", templatesHandler)


			picturesDir := http.Dir("../pictures")
			picturesHandler := http.FileServer(picturesDir)
			picturesHandler = http.StripPrefix("/pictures/", picturesHandler)
			http.Handle("/pictures/", picturesHandler)
	*/
	http.ListenAndServe(":8080", nil)
}
