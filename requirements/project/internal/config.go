package internal

import "html/template"

const (
	INDEX_TEMPLATE_PATH = "./templates/index.html"
	BANNERS_PATH        = "./banners/"
	MAXINPUTLENGTH      = 200
)

var Font = make(map[rune][]string, 95)
var AsciiArtTmpl = template.Must(template.ParseFiles("./templates/asciiArt.html"))

type Art struct {
	Title string
	Lines string
}

const (
	CharacterHeight = 8
	STANDARD_HASH   = "ac85e83127e49ec42487f272d9b9db8b"
	SHADOW_HASH     = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	THINKERTOY_HASH = "86d9947457f6a41a18cb98427e314ff8"
)
