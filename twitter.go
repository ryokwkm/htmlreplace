package html_replace

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
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
	ret = AdjustBrackets(ret)
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

func DelStrings(body string, dels []string) string {
	for _, del := range dels {
		body = strings.Replace(body, del, "", -1)
	}
	return body
}

func GetMention(body string) string {
	r := regexp.MustCompile(`@(w*[一-龠_ぁ-ん_ァ-ヴーａ-ｚＡ-Ｚa-zA-Z0-9]+|[a-zA-Z0-9_]+|[a-zA-Z0-9_]w*)`)
	l := r.FindAllStringSubmatch(body, -1)
	mentions := []string{}
	if l != nil {
		for _, e := range l {
			mentions = append(mentions, e[0])
		}
	}
	return strings.Join(mentions, " ")
}

func AdjustBrackets(str string) string {
	type Pair struct {
		Front string
		Back  string
	}

	pairs := []Pair{
		{Front: "【", Back: "】"},
		{Front: "「", Back: "」"},
		{Front: "『", Back: "』"},
		{Front: "（", Back: "）"},
		{Front: "（", Back: "）"},
		{Front: "［", Back: "］"},
		{Front: "＜", Back: "＞"},
	}
	for _, p := range pairs {
		fc := strings.Count(str, p.Front)
		bc := strings.Count(str, p.Back)
		if fc != bc {
			str = strings.Replace(str, p.Front, "", -1)
			str = strings.Replace(str, p.Back, "", -1)
		}
	}
	//。を削除
	str = strings.Replace(str, "。", "", -1)
	return str
}

//ランダムに取得。同じ言葉を繰り返すと面白くなる
func ChoiceRandom(words []string, limit int) []string {
	if len(words) < limit {
		limit = len(words)
	}

	ret := []string{}
	for i := 0; i < limit; i++ {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(words))
		ret = append(ret, words[i])
	}
	return ret
}

//本文にURLをつける
func ConvertURL(tweetWord string, url string) string {
	const TweetLimit = 125
	//１０文字以上余白があるならつぶやく
	if len([]rune(tweetWord)) > TweetLimit {
		remainder := len([]rune(tweetWord)) - TweetLimit
		tweetWord = string([]rune(tweetWord)[:remainder])
	}
	return tweetWord + url
}
