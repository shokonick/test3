package utils

import (
	"github.com/gocolly/colly"
	"strings"
    "github.com/google/uuid"
	"github.com/google/go-querystring/query"
	"github.com/OwO-Network/gdeeplx"
	"os"
	"errors"
)

type LangOut struct {
	OutputText string
	AutoDetect string
}

func TranslateGoogle(to string, from string, text string) (LangOut, error) {
	// For some reason google uses no for norwegian instead of nb like the rest of the translators. This is for the All function primarily
	if to == "nb" {
		to = "no"
	} else if from == "nb" {
		to = "no"
	}
	var ToValid bool
	var FromValid bool
	for _, v := range LangListGoogle("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target Language Code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}

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
		To   string `url:"tl"`
		UI   string `url:"hl"`
		From string `url:"sl"`
		Text string `url:"q"`
	}
	opt := Options{to, to, from, text}
	v, _ := query.Values(opt)
	url := "https://translate.google.com/m?" + v.Encode()
	sc.Visit(url)
	var langout LangOut
	langout.OutputText = answer
	return langout, nil
}
func TranslateReverso(to string, from string, query string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListReverso("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	json := []byte(`{ "format": "text", "from": "` + from + `", "to": "` + to + `", "input":"` + query + `", "options": {"sentenceSplitter": false, "origin":"translation.web", contextResults: false, languageDetection: true} }`)
	reversoOut := PostRequest("https://api.reverso.net/translate/v1/translation", json)
	gjsonArr := reversoOut.Get("translation").Array()
	var langout LangOut
	langout.OutputText = gjsonArr[0].String()
	return langout, nil
}
func TranslateLibreTranslate(to string, from string, query string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListLibreTranslate("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	json := []byte(`{"q":"` + query + `","source":"` + from + `","target":"` + to + `"}`)
	// TODO: Make it configurable
	libreTranslateOut := PostRequest("https://translate.argosopentech.com/translate", json)
	gjsonArr := libreTranslateOut.Get("translatedText").Array()
	var langout LangOut
	langout.OutputText = gjsonArr[0].String()
	return langout, nil
}
func TranslateWatson(to string, from string, query string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListWatson("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	json := []byte(`{"text":"` + query + `","source":"` + from + `","target":"` + to + `"}`)
	watsonOut := PostRequest("https://www.ibm.com/demos/live/watson-language-translator/api/translate/text", json)
	gjsonArr := watsonOut.Get("payload.translations.0.translation").Array()
	var langout LangOut
	langout.OutputText = gjsonArr[0].String()
	return langout, nil
}
func TranslateMyMemory(to string, from string, text string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListMyMemory("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	type Options struct {
		Translate string `url:"langpair"`
		Text      string `url:"q"`
	}
	opt := Options{ from+"|"+to, text }
	v, _ := query.Values(opt)
	myMemoryOut := GetRequest("https://api.mymemory.translated.net/get?" + v.Encode())
	gjsonArr := myMemoryOut.Get("responseData.translatedText").Array()
	var langout LangOut
	langout.OutputText = gjsonArr[0].String()
	return langout, nil
}
func TranslateYandex(to string, from string, text string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListYandex("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	type Options struct {
		Translate string `url:"lang"`
		Text      string `url:"text"`
		Srv       string `url:"srv"`
		Id        string `url:"sid"`
	}
    uuidWithHyphen := uuid.New()
    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	opt := Options{from + "-" + to, text, "android", uuid+"-0-0"}
	v, _ := query.Values(opt)

	yandexOut := PostRequest("https://translate.yandex.net/api/v1/tr.json/translate?" + v.Encode(), []byte(""))
	gjsonArr := yandexOut.Get("text.0").Array()
	var langout LangOut
	langout.OutputText = gjsonArr[0].String()
	return langout, nil
}
func TranslateDeepl(to string, from string, text string) (LangOut, error) {
	var ToValid bool
	var FromValid bool
	for _, v := range LangListDeepl("sl") {
		if v.Id == to {
			ToValid = true	
		}
		if v.Id == from {
			FromValid = true
		}
		if FromValid == true && ToValid == true {
			break
		}
	}
	if ToValid != true {
		return LangOut{}, errors.New("Target language code invalid")
	}
	if FromValid != true {
		return LangOut{}, errors.New("Source language code invalid")
	}
	answer, err := gdeeplx.Translate(text, from, to, 0)
	if err != nil {
		return LangOut{}, errors.New("failed")
	}
	answer1 := answer.(map[string]interface{})
	ans:= answer1["data"].(string)
	var langout LangOut
	langout.OutputText = ans
	return langout, nil
}
func TranslateAll(to string, from string, query string) (string, string, string, string, string, string, string) {
	reverso, _ := TranslateReverso(to, from, query)
	google, _ := TranslateGoogle(to, from, query)
	libretranslate, _ := TranslateLibreTranslate(to, from, query)
	watson, _ := TranslateWatson(to, from, query)
	mymemory, _ := TranslateMyMemory(to, from, query)
	yandex, _ := TranslateYandex(to, from, query)
	deepl, _ := TranslateDeepl(to, from, query)
	return google.OutputText, reverso.OutputText, libretranslate.OutputText, watson.OutputText, mymemory.OutputText, yandex.OutputText, deepl.OutputText
}
