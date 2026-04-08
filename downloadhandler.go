package main

import "net/http"

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
    w.Header().Set("Content-Type", "text/plain")
	w.Write(indexPageFillings.Art)
}
