package internal

import (
	"ascii-art-web/error"
	"html/template"
	"net/http"
	"os"
)

const maxCommentLength = 200

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
	if len(inputText) > maxCommentLength {
		error.BadRequest(w, r)
		return
	}
	banner := r.FormValue("banner")
	if !isBanner(banner) {
		error.BadRequest(w, r)
		return
	}
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

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	filePath, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			error.NotFoundError(w, r)
			return
		}
		error.InternalServerError(w, r)
		return
	}
	if filePath.IsDir() {
		error.NotFoundError(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func isBanner(banner string) bool { // Avoid an internal server error crash
	banners := []string{"standard", "shadow", "thinkertoy"}
	for _, v := range banners {
		if v == banner {
			return true
		}
	}
	return false
}
