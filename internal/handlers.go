package internal

import (
	"fmt"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method: 1", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	inputText := r.FormValue("inputText")
	banner := r.FormValue("banner")
	if inputText == "" {
		http.Error(w, "Input text is required", http.StatusBadRequest)
		return
	}
	var toPrint Art
	err2 := ParsFile("./banners/", banner)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
	toPrint = StringAscii(inputText)
	for _, c := range toPrint.Lines {
		fmt.Fprintf(w, "%s", c)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method: 2", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./templates/index.html")
}
