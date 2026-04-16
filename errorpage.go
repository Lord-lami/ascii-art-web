package main

import (
	"log"
	"net/http"
	"strings"
)

// errorPage shows a page that redirects to the main page in 5 seconds 
// with the passed parameters slotted in their necessary places
func errorPage(w http.ResponseWriter, statusCode int, title, message string) {
	if statusCode < 400 {
		log.Println(statusCode, "is not an error code")
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")
	layout.Meta = `<meta http-equiv="refresh" content="5;url=/">`
	layout.Title = title
	var pageContent strings.Builder
	templates.ExecuteTemplate(&pageContent, "error.html", message)
	layout.Content = pageContent.String()
	templates.ExecuteTemplate(w, "layout.html", layout)
}
