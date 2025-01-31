package internal

const INDEX_TEMPLATE_PATH = "./templates/index.html"
const BANNERS_PATH = "./banners/"

var Font = make(map[rune][]string, 95)

type Art struct {
	Title string
	Lines []string
}

const CharacterHeight = 8
const STANDARD_HASH = "ac85e83127e49ec42487f272d9b9db8b"
const SHADOW_HASH = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
const THINKERTOY_HASH = "86d9947457f6a41a18cb98427e314ff8"
