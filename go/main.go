package main

import (
	"net/http"
)

func main() {
	mainDir := http.Dir("../static")
	staticHandler := http.FileServer(mainDir)
	http.Handle("/", staticHandler)

	stylesDir := http.Dir("../styles")
	stylesHandler := http.FileServer(stylesDir)
	staticHandler = http.StripPrefix("/styles/", stylesHandler)
	http.Handle("/styles/", staticHandler)

	/*
		picturesDir := http.Dir("../pictures")
		picturesHandler := http.FileServer(picturesDir)
		picturesHandler = http.StripPrefix("/pictures/", picturesHandler)
		http.Handle("/pictures/", picturesHandler)
	*/
	http.ListenAndServe(":8080", nil)
}
