package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/page.html",
	"templates/index.html", "templates/error.html"))

var indexPageFillings struct {
	Text           string
	Banners        []string
	Selected       int
	Art            []byte
	DownloadButton string
}

var page struct {
	Meta    string
	Title   string
	Content string
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	page.Meta = ""
	page.Title = "ASCII Art"
	var pageContent strings.Builder
	templates.ExecuteTemplate(&pageContent, "index.html", indexPageFillings)
	page.Content = pageContent.String()

	templates.ExecuteTemplate(w, "page.html", page)
}

// loadBanners loads the banner files
// in the banner folder into the
// indexPageFillings.Banners slice
func loadBanners() {
	bannerDir, _ := os.ReadDir("banners")
	for _, banner := range bannerDir {
		bannerName := strings.TrimSuffix(banner.Name(), ".txt")
		indexPageFillings.Banners = append(indexPageFillings.Banners, bannerName)
	}
	indexPageFillings.Selected = slices.Index(indexPageFillings.Banners, "standard")
}

func invalidPathHandler(w http.ResponseWriter, r *http.Request) {
	errorPage(w, http.StatusNotFound, "Not Found", "Incorrect URL: This page doesn't exist")
}

func main() {
	loadBanners()
	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	http.HandleFunc("GET /export/{$}", downloadHandler)

	http.HandleFunc("/", invalidPathHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
