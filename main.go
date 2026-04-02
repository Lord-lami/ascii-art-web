package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/ascii-art-web.html",
	"templates/index.html"))

var fillings struct {
	Text     string
	Banners  []string
	Selected int
	Art      []byte
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", fillings)
}

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	fillings.Text = r.FormValue("text")
	banner := r.FormValue("banner")
	fillings.Selected = slices.Index(fillings.Banners, banner)
	if fillings.Selected == -1 {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}
	var err error
	fillings.Art, err = exec.Command("./ascii-art-web", fillings.Text, banner).CombinedOutput()
	if err != nil {
		log.Println(string(fillings.Art))
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func invalidPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusNotFound)
}

func main() {
	// Load the banners folder into the
	// html banner select options
	bannerDir, _ := os.ReadDir("banners")
	for _, banner := range bannerDir {
		bannerName := strings.TrimSuffix(banner.Name(), ".txt")
		fillings.Banners = append(fillings.Banners, bannerName)
	}
	fillings.Selected = slices.Index(fillings.Banners, "standard")

	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	http.HandleFunc("/", invalidPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
