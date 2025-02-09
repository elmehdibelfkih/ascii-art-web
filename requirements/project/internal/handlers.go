package internal

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ascii-art-web/error"
)

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

	inputText := r.FormValue("before-generate")
	if len(inputText) > MAXINPUTLENGTH {
		error.BadRequest(w, r)
		return
	}
	banner := r.FormValue("banner")
	if !isBanner(banner) || inputText == "" || banner == "" {
		error.BadRequest(w, r)
		return
	}
	var toPrint Art
	err2 := ParsFile(BANNERS_PATH, banner)
	if err2 != nil {
		error.InternalServerError(w, r)
		return
	}
	toPrint = StringAscii(inputText)
	err3 := ASCIIARTTMPL.Execute(w, toPrint)
	if err3 != nil {
		error.InternalServerError(w, r)
		return
	}
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

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	before_generate := r.FormValue("before-generate")
	after_generate := r.FormValue("after-generate")
	banner := r.FormValue("banner")

	var asciiStruct Art
	var ascii string

	err := r.ParseForm()
	if err != nil {
		error.InternalServerError(w, r)
		return
	}

	if len(before_generate) > MAXINPUTLENGTH {
		error.BadRequest(w, r)
		return
	}

	if len(before_generate) == 0 && len(after_generate) == 0 {
		error.BadRequest(w, r)
		return
	}

	if len(after_generate) == 0 {
		ParsFile(BANNERS_PATH, banner)
		if !isBanner(banner) {
			error.BadRequest(w, r)
		}
		asciiStruct = StringAscii(before_generate)
		ascii = asciiStruct.Lines
	} else {
		ascii = after_generate
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(ascii)))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	fmt.Fprint(w, ascii)
}
