package main

import (
	"log"
	"net/http"
	"os"
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

// loadBanners loads the banner files
// in the banner folder into the
// fillings.Banners slice
func loadBanners() {
	bannerDir, _ := os.ReadDir("banners")
	for _, banner := range bannerDir {
		bannerName := strings.TrimSuffix(banner.Name(), ".txt")
		fillings.Banners = append(fillings.Banners, bannerName)
	}
	fillings.Selected = slices.Index(fillings.Banners, "standard")
}

func invalidPathHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusNotFound)
}

func main() {
	loadBanners()
	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	http.HandleFunc("/", invalidPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
