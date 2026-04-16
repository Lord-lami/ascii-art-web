package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"text/template"
)

// Embedding
//
//go:embed templates/*.html
var templatesFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

var indexPageFillings struct {
	Text                string
	ColorDetailsChanged bool
	Color               string
	ColoredText         string
	Alignment           string
	Banners             []string
	Selected            int
	Art                 []byte
	DownloadButton      string
}

var layout struct {
	Meta    string
	Title   string
	Content string
}

// mainPageHandler handles the main page of the website
// it populates the respective pipelines in the templates
// with indePageFillings and puts the main page in the layout
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	layout.Meta = ""
	layout.Title = "ASCII Art"
	var pageContent strings.Builder
	templates.ExecuteTemplate(&pageContent, "index.html", indexPageFillings)
	layout.Content = pageContent.String()

	templates.ExecuteTemplate(w, "layout.html", layout)
}

// loadDefaults loads the banner files
// in the banner folder into the
// indexPageFillings.Banners slice
// and sets the default color for
// indexPageFillings.Color
func loadDefaults() {
	bannerDir, _ := os.ReadDir("banners")
	for _, banner := range bannerDir {
		bannerName := strings.TrimSuffix(banner.Name(), ".txt")
		indexPageFillings.Banners = append(indexPageFillings.Banners, bannerName)
	}
	indexPageFillings.Selected = slices.Index(indexPageFillings.Banners, "standard")
	indexPageFillings.Color = "#ffffff"
}

func invalidPathHandler(w http.ResponseWriter, r *http.Request) {
	errorPage(w, http.StatusNotFound, "Not Found", "Incorrect URL: This page doesn't exist")
}

func main() {
	loadDefaults()
	http.HandleFunc("GET /{$}", mainPageHandler)

	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	http.HandleFunc("POST /ascii-art/{$}", asciiArtPageHandler)

	http.HandleFunc("GET /export/{$}", downloadHandler)

	http.HandleFunc("/", invalidPathHandler)
	log.Println("Server running on port 8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
