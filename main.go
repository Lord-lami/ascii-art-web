package main

import (
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/ascii-art.html"))

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	asciiArt, err := exec.Command("./ascii-art-web", text, banner).CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = templates.ExecuteTemplate(w, "ascii-art.html", asciiArt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
