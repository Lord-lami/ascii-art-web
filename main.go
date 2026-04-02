package main

import (
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/ascii-art.html",
	"templates/index.html"))

var fillings struct {
	Text   string
	Select struct {
		Standard   string
		Shadow     string
		Thinkertoy string
	}
	Art []byte
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", fillings)
}

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	asciiArt, err := exec.Command("./ascii-art-web", text, banner).CombinedOutput()
	if err != nil {
		log.Println(string(asciiArt))
		fillings.Art = asciiArt
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	fillings.Text = text
	selected := "selected"
	switch banner {
	case "standard":
		fillings.Select.Standard = selected
	case "shadow":
		fillings.Select.Shadow = selected
	case "thinkertoy":
		fillings.Select.Thinkertoy = selected
	}
	fillings.Art = asciiArt
	http.Redirect(w, r, "/", http.StatusFound)
}

func invalidPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusNotFound)
}

func main() {
	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	http.HandleFunc("/", invalidPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
