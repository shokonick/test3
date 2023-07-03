package utils

import (
	"github.com/gocolly/colly"
	"os"
)

func TranslateGoogle(to string, from string, query string) string {
	UserAgent, ok := os.LookupEnv("SIMPLYTRANSLATE_USER_AGENT")
	if !ok {
		UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	}
	sc := colly.NewCollector(colly.AllowedDomains("translate.google.com"), colly.UserAgent(UserAgent))
	var answer string
	sc.OnHTML("div.result-container", func(e *colly.HTMLElement) {
		answer = e.Text
	})
	url := "https://translate.google.com/m?sl=" + from + "&tl=" + to + "&hl=" + to + "&q=" + query
	sc.Visit(url)
	return answer
}
func TranslateReverso(to string, from string, query string) string {
	json := []byte(`{ "format": "text", "from": "` + from + `", "to": "` + to + `", "input":"` + query + `", "options": {"sentenceSplitter": false, "origin":"translation.web", contextResults: false, languageDetection: true} }`)
	reversoOut := GetRequest("https://api.reverso.net/translate/v1/translation", json)
	gjsonArr := reversoOut.Get("translation").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateLibreTranslate(to string, from string, query string) string {
	json := []byte(`{"q":"`+query+`","source":"`+from+`","target":"`+to+`"}`)
	// TODO: Make it configurable
	libreTranslateOut := GetRequest("https://translate.argosopentech.com/translate", json)
	gjsonArr := libreTranslateOut.Get("translatedText").Array()
	answer := gjsonArr[0].String()
	return answer
}
func TranslateAll(to string, from string, query string) string {
	reverso := TranslateReverso(to, from, query)
	google := TranslateGoogle(to, from, query)
	libretranslate := TranslateLibreTranslate(to, from, query)
	return "Google: " + google + "\nReverso: " + reverso + "\nLibreTranslate: " + libretranslate
}
