package main

import (
	"log"
	"net/http"
	"os/exec"
	"slices"
)

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	indexPageFillings.Text = r.FormValue("text")
	banner := r.FormValue("banner")
	indexPageFillings.Selected = slices.Index(indexPageFillings.Banners, banner)

	if indexPageFillings.Selected == -1 || indexPageFillings.Text == "" {
		log.Panicln("Banner", banner, "is not in the banners folder")
		errorPage(w, http.StatusNotFound, "Banner Not Found",
			"There is no "+banner+" banner font.")
		return
	}

	var err error
	indexPageFillings.Art, err = exec.Command("./ascii-art-web", indexPageFillings.Text, banner).CombinedOutput()

	if err != nil {
		log.Println(string(indexPageFillings.Art))
		errorPage(w, http.StatusInternalServerError, "Invalid Input",
			"None ASCII Characters are Invalid")
		return
	}
	indexPageFillings.DownloadButton = `<form action="/export/" method="GET">
	<button type="submit"><strong>📥 Download as text file</strong></button>
	</form>`
	http.Redirect(w, r, "/", http.StatusFound)
}
