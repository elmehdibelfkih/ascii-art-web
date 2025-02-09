package internal

import (
	"ascii-art-web/error"
	"html/template"
	"log"
)

func isBanner(banner string) bool { // Avoid an internal server error crash
	banners := []string{"standard", "shadow", "thinkertoy"}
	for _, v := range banners {
		if v == banner {
			return true
		}
	}
	return false
}

func CreateTemplates() {
	tmpl, err := template.ParseFiles("./templates/asciiArt.html")
	if err != nil {
		log.Fatal(err)
	}
	ASCIIARTTMPL = template.Must(tmpl, err)
	tmpl1, err1 := template.ParseFiles("./templates/error.html")
	if err1 != nil {
		log.Fatal(err1)
	}
	error.ERRORTMPL = template.Must(tmpl1, err1)
}
