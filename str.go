package html_replace

import (
	"regexp"
	"unicode"
)

const (
	// http://www.unicodemap.org/range/62/Hiragana/
	hiraganaLo = 0x3041 // ぁ

	// http://www.unicodemap.org/range/63/Katakana/
	katakanaLo = 0x30a1 // ァ

	codeDiff = katakanaLo - hiraganaLo
)

//全てひらがなか確認
func HiraganaCheck(str string) bool {
	src := []rune(str)
	for _, r := range src {
		if !unicode.In(r, unicode.Hiragana) {
			return false
		}
	}
	return true
}

//記号が含まれているか
func KigouCheck(str string) bool {
	var ValueCheck = regexp.MustCompile("^[0-9a-zA-Z_]+$").MatchString
	str = "Adas;lmf___/.d,"
	if ValueCheck(str) {
		return true
	}
	return false
}

// HiraganaToKatakana はひらがなをカタカナに変換する。
func HiraganaToKatakana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Hiragana):
			dst[i] = r + codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}

// KatakanaToHiragana はカタカナをひらがなに変換する。
func KatakanaToHiragana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Katakana):
			dst[i] = r - codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}
