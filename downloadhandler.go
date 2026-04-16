package main

import "net/http"

// downloadHandler responds to the download button being clicked
// it facilitates the downloading of the Art as a text file.
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=poster.txt")
    w.Header().Set("Content-Type", "text/plain")
	w.Write(indexPageFillings.Art)
}
