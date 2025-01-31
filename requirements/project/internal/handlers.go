package internal

import (
	"ascii-art-web/error"
	"html/template"
	"net/http"
)

var AsciiArtTmpl = template.Must(template.ParseFiles("./templates/asciiArt.html"))

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		error.MethodNotAllowed(w, r)
		return
	}
	err := r.ParseForm()
	if err != nil {
		error.InternalServerError(w, r)
		return
	}
	inputText := r.FormValue("inputText")
	banner := r.FormValue("banner")
	if inputText == "" || banner == "" {
		error.NotContainTheRequiredData(w, r)
		return
	}
	var toPrint Art
	err2 := ParsFile(BANNERS_PATH, banner)
	if err2 != nil {
		error.InternalServerError(w, r)
		return
	}
	toPrint = StringAscii(inputText)
	AsciiArtTmpl.Execute(w, toPrint)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		error.NotFoundError(w, r)
		return
	}
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	http.ServeFile(w, r, INDEX_TEMPLATE_PATH)
}
