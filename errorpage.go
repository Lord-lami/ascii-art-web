package main

import (
	"log"
	"net/http"
	"strings"
)

func errorPage(w http.ResponseWriter, statusCode int, title, message string) {
	if statusCode < 400 {
		log.Println(statusCode, "is not an error code")
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")
	page.Meta = `<meta http-equiv="refresh" content="5;url=/">`
	page.Title = title
	var pageContent strings.Builder
	templates.ExecuteTemplate(&pageContent, "error.html", message)
	*page.Content = pageContent.String()
	templates.ExecuteTemplate(w, "page.html", page)
}
