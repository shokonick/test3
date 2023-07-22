package utils

import (
	"github.com/gocolly/colly"
	"github.com/google/go-querystring/query"
	"os"
)

func TranslateGoogle(to string, from string, text string) string {
	UserAgent, ok := os.LookupEnv("SIMPLYTRANSLATE_USER_AGENT")
	if !ok {
		UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	}
	sc := colly.NewCollector(colly.AllowedDomains("translate.google.com"), colly.UserAgent(UserAgent))
	var answer string
	sc.OnHTML("div.result-container", func(e *colly.HTMLElement) {
		answer = e.Text
	})
	type Options struct {
	  To string `url:"tl"`
	  UI string `url:"hl"`
	  From string `url:"sl"`
	  Text string `url:"q"`
	}
	opt := Options{ to, to, from, text }
	v, _ := query.Values(opt)
	url := "https://translate.google.com/m?"+v.Encode()
	sc.Visit(url)
	return answer
}
func TranslateReverso(to string, from string, query string) string {
	json := []byte(`{ "format": "text", "from": "` + from + `", "to": "` + to + `", "input":"` + query + `", "options": {"sentenceSplitter": false, "origin":"translation.web", contextResults: false, languageDetection: true} }`)
	reversoOut := PostRequest("https://api.reverso.net/translate/v1/translation", json)
	gjsonArr := reversoOut.Get("translation").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateLibreTranslate(to string, from string, query string) string {
	json := []byte(`{"q":"`+query+`","source":"`+from+`","target":"`+to+`"}`)
	// TODO: Make it configurable
	libreTranslateOut := PostRequest("https://translate.argosopentech.com/translate", json)
	gjsonArr := libreTranslateOut.Get("translatedText").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateWatson(to string, from string, query string) string {
	json := []byte(`{"text":"`+query+`","source":"`+from+`","target":"`+to+`"}`)
	watsonOut:= PostRequest("https://www.ibm.com/demos/live/watson-language-translator/api/translate/text", json)
	gjsonArr := watsonOut.Get("payload.translations.0.translation").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateYandex(to string, from string, text string) string {
	type Options struct {
	  Translate string `url:"lang"`
	  Text string `url:"text"`
	  Srv string `url:"srv"`
	  Id string `url:"id"`
	  Reason string `url:"reason"`
	}
	opt := Options{ from+"-"+to, text, "tr-mobile", "c2317111.64bac36a.ab16ef22.74722d6d6f62696c65-0-0", "submit"}
	v, _ := query.Values(opt)

	yandexOut := GetRequest("https://translate.yandex.net/api/v1/tr.json/translate?"+v.Encode())
	gjsonArr := yandexOut.Get("text.0").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateAll(to string, from string, query string) string {
	reverso := TranslateReverso(to, from, query)
	google := TranslateGoogle(to, from, query)
	libretranslate := TranslateLibreTranslate(to, from, query)
	return "Google: " + google + "\nReverso: " + reverso + "\nLibreTranslate: " + libretranslate
}
