package html_replace

import (
	"regexp"
	"strings"
)

func GetOnlyTexts(list []string) []string {
	ret := []string{}
	for _, body := range list {
		ret = append(ret, GetOnlyText(body))
	}
	return ret
}

func GetOnlyText(body string) string {
	ret := DelMention(body)
	ret = DelHash(ret)
	ret = DelURL(ret)
	return ret
}

func DelMention(body string) string {
	rep := regexp.MustCompile(`@(w*[一-龠_ぁ-ん_ァ-ヴーａ-ｚＡ-Ｚa-zA-Z0-9]+|[a-zA-Z0-9_]+|[a-zA-Z0-9_]w*)`)
	ret := rep.ReplaceAllString(body, "")
	ret = strings.TrimSpace(ret)
	return ret
}

func DelHash(body string) string {
	rep := regexp.MustCompile(`[#＃][Ａ-Ｚａ-ｚA-Za-z一-鿆0-9０-９ぁ-ヶｦ-ﾟー]+`)
	ret := rep.ReplaceAllString(body, "")
	ret = strings.TrimSpace(ret)
	return ret
}

func DelURL(body string) string {
	rep := regexp.MustCompile(`https?://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?`)
	ret := rep.ReplaceAllString(body, "")
	ret = strings.TrimSpace(ret)
	return ret
}
