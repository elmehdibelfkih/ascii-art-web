package internal

import "html/template"

const ( //
	CharacterHeight     = 8
	MAXINPUTLENGTH      = 200
	STANDARD_HASH       = "ac85e83127e49ec42487f272d9b9db8b"
	SHADOW_HASH         = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	THINKERTOY_HASH     = "86d9947457f6a41a18cb98427e314ff8"
	INDEX_TEMPLATE_PATH = "./templates/index.html"
	BANNERS_PATH        = "./banners/"
)

var FONT = make(map[rune][]string, 95)
var ASCIIARTTMPL *template.Template

type Art struct {
	Title string
	Lines string
}
