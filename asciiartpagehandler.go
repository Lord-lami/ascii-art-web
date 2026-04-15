package main

import (
	"log"
	"net/http"
	"os/exec"
	"slices"
	"strings"
)

func asciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	indexPageFillings.Text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\n")
	if indexPageFillings.Text == "" {
		log.Println("Someone bypassed the required field, textarea")
		errorPage(w, http.StatusBadRequest, "No Text",
			"There is no text to draw. Write something")
		return
	}

	prevColor := indexPageFillings.Color
	indexPageFillings.Color = r.FormValue("color")
	prevColoredText := indexPageFillings.ColoredText
	indexPageFillings.ColoredText = strings.ReplaceAll(r.FormValue("colored-text"), "\r\n", "\n")
	indexPageFillings.ColorDetailsChanged = prevColor != indexPageFillings.Color || prevColoredText != indexPageFillings.ColoredText

	indexPageFillings.Alignment = r.FormValue("alignment")
	switch indexPageFillings.Alignment {
	case "", "left", "right", "center", "justify":
	default:
		log.Println("Alignment", indexPageFillings.Alignment, "is not a valid alignment")
		errorPage(w, http.StatusBadRequest, "Invalid Alignment",
			indexPageFillings.Alignment+" is an invalid alignment")
		return
	}

	banner := r.FormValue("banner")
	indexPageFillings.Selected = slices.Index(indexPageFillings.Banners, banner)

	if indexPageFillings.Selected == -1 {
		log.Println("Banner", banner, "is not in the banners folder")
		errorPage(w, http.StatusNotFound, "Banner Not Found",
			"There is no "+banner+" banner font.")
		return
	}

	var err error
	indexPageFillings.Art, err = exec.Command("./ascii-art-full",
		"--align="+indexPageFillings.Alignment,
		"--color="+indexPageFillings.Color,
		indexPageFillings.ColoredText,
		indexPageFillings.Text,
		banner).CombinedOutput()

	if err != nil {
		// Check that the log prints an invalid character error
		// If it is a different type of error a new input validation
		// must be added.
		log.Println(string(indexPageFillings.Art))
		errorPage(w, http.StatusInternalServerError, "Invalid Input",
			"None ASCII Characters are Invalid")
		return
	}
	indexPageFillings.DownloadButton = `<form action="/export/" method="GET">
	<button type="submit"><strong>📥 Download as a text file</strong></button>
	</form>`
	http.Redirect(w, r, "/", http.StatusFound)
}
