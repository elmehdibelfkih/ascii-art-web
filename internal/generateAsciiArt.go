package internal

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"strings"
)

var Font = make(map[rune][]string, 95)

type Art struct {
	Lines []string
}

const CharacterHeight = 8
const STANDARD_HASH = "ac85e83127e49ec42487f272d9b9db8b"
const SHADOW_HASH = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
const THINKERTOY_HASH = "86d9947457f6a41a18cb98427e314ff8"

func ParsFile(fontFileName string, font string) error {
	data, err := os.ReadFile(fontFileName + font + ".txt")
	if err != nil {
		return err
	}
	hashObj := md5.New()
	_, err1 := hashObj.Write(data)
	if err1 != nil {
		return err1
	}
	hash := fmt.Sprintf("%x", hashObj.Sum(nil))

	if font == "standard" || font == "shadow" {
		if (font == "standard" && hash != STANDARD_HASH) || (font == "shadow" && hash != SHADOW_HASH) {
			return errors.New("hash does not match")
		}
		str := strings.Trim(string(data), "\n")
		tmp := strings.Split(str, "\n\n")
		for i, j := 32, 0; i <= 126; i++ {
			Font[rune(i)] = strings.Split(tmp[j], "\n")
			j++
		}
	} else if font == "thinkertoy" {
		if hash != THINKERTOY_HASH {
			return errors.New("hash does not match")
		}
		str := strings.Trim(string(data), "\r\n")
		tmp := strings.Split(str, "\r\n\r\n")
		for i, j := 32, 0; i <= 126; i++ {
			Font[rune(i)] = strings.Split(tmp[j], "\r\n")
			j++
		}
	}
	return nil
}

func StringAscii(str string) Art {
	ret := Art{}

	if strings.ReplaceAll(str, `\n`, "") == "" {
		counter := strings.Count(str, `\n`)
		fmt.Print(strings.Repeat("\n", counter))
		for i := 0; i < counter; i++ {
			ret.Lines = append(ret.Lines, "\n")
		}
		return ret
	}
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < CharacterHeight; i++ {
				tmpLine := ""
				for _, c := range line {
					if c < 127 && c > 31 && i < len(Font[c]) {
						tmpLine += Font[c][i]
					}
				}
				ret.Lines = append(ret.Lines, tmpLine+"\n")
			}
		} else {
			ret.Lines = append(ret.Lines, "\n")
		}
	}
	return ret
}
