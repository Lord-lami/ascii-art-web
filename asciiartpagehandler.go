package main

import (
	"log"
	"net/http"
	"os/exec"
	"slices"
)

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	fillings.Text = r.FormValue("text")
	banner := r.FormValue("banner")
	fillings.Selected = slices.Index(fillings.Banners, banner)
	if fillings.Selected == -1 || fillings.Text == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	var err error
	fillings.Art, err = exec.Command("./ascii-art-web", fillings.Text, banner).CombinedOutput()
	if err != nil {
		log.Println(string(fillings.Art))
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
