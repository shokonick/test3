package utils

import (
	"github.com/gocolly/colly"
	"os"
	"strings"
)

type List struct {
	Name string
	Id   string
}

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
func LangListGoogle(listType string) []List {
	UserAgent, ok := os.LookupEnv("SIMPLYTRANSLATE_USER_AGENT")
	if !ok {
		UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	}
	sc := colly.NewCollector(colly.AllowedDomains("translate.google.com"), colly.UserAgent(UserAgent))
	var ListData []List
	sc.OnHTML("div.root-container", func(e *colly.HTMLElement) {
		e.ForEach("div.language-item", func(i int, el *colly.HTMLElement) {
			var ListInfo List
			ListInfo.Name = el.ChildText("a")
			if listType == "tl" {
				ListInfo.Id = strings.TrimPrefix(strings.TrimSuffix(el.ChildAttr("a", "href"), "&hl=en-US"), "./m?sl=en&tl=")
			} else {
				ListInfo.Id = strings.TrimPrefix(strings.TrimSuffix(el.ChildAttr("a", "href"), "&tl=en&hl=en-US"), "./m?sl=")
			}
			ListData = append(ListData, ListInfo)
		})
	})
	url := "https://translate.google.com/m?sl=en&tl=en&hl=en-US&mui=" + listType
	sc.Visit(url)
	return ListData
}
func LangListLibreTranslate(listType string) []List {
	// TODO: Make it configurable
	libreTranslateOut := GetRequest("https://translate.argosopentech.com/languages", []byte(""))
	gjsonArr := libreTranslateOut.Array()
	var ListData []List
	for _, r := range gjsonArr {
		code := r.Get("code").String()
		name := r.Get("name").String()
	
		ListData = append(ListData, List{Id: code, Name: name})
	}
	if listType == "sl" {
		auto := List{
			Id:   "auto",
			Name: "Detect Language",
		}
		ListData = append(ListData, auto)
	}
	return ListData
}
func LangListReverso(listType string) []List {
	// IDs got from original simplytranslate-web and trial and error. Usually first three letters of language.
	var ListData = []List{
		List{
			Id:   "ara",
			Name: "Arabic",
		},
		List{
			Id:   "chi",
			Name: "Chinese (Simplified)",
		},
		List{
			Id:   "cze",
			Name: "Czech",
		},
		List{
			Id:   "dan",
			Name: "Danish",
		},
		List{
			Id:   "dut",
			Name: "Dutch",
		},
		List{
			Id:   "eng",
			Name: "English",
		},
		List{
			Id:   "fra",
			Name: "French",
		},
		List{
			Id:   "ger",
			Name: "German",
		},
		List{
			Id:   "gre",
			Name: "Greek",
		},
		List{
			Id:   "Hebrew",
			Name: "heb",
		},
		List{
			Id:   "hin",
			Name: "Hindi",
		},
		List{
			Id:   "hun",
			Name: "Hungarian",
		},
		List{
			Id:   "ita",
			Name: "Italian",
		},
		List{
			Id:   "jpn",
			Name: "Japanese",
		},
		List{
			Id:   "kor",
			Name: "Korean",
		},
		List{
			Id:   "per",
			Name: "Persian",
		},
		List{
			Id:   "pol",
			Name: "Polish",
		},
		List{
			Id:   "por",
			Name: "Portugese",
		},
		List{
			Id:   "rum",
			Name: "Romanian",
		},
		List{
			Id:   "rus",
			Name: "Russian",
		},
		List{
			Id:   "slo",
			Name: "Slovakian",
		},
		List{
			Id:   "spa",
			Name: "Spanish",
		},
		List{
			Id:   "swe",
			Name: "Swedish",
		},
		List{
			Id:   "tha",
			Name: "Thai",
		},
		List{
			Id:   "tur",
			Name: "Turkish",
		},
		List{
			Id:   "ukr",
			Name: "Ukrainian",
		},
	}
	return ListData
}
func LangListDeepl(listType string) []List {
	// IDs got from deepl.com/translator
	// Every time you change language, the # will get updated with the lang code.
	var ListData = []List{
		List{
			Id:   "bg",
			Name: "Bulgarian",
		},
		List{
			Id:   "zh",
			Name: "Chinese",
		},
		List{
			Id:   "cs",
			Name: "Czech",
		},
		List{
			Id:   "da",
			Name: "Danish",
		},
		List{
			Id:   "nl",
			Name: "Dutch",
		},
		List{
			Id:   "en",
			Name: "English",
		},
		List{
			Id:   "et",
			Name: "Estonian",
		},
		List{
			Id:   "fi",
			Name: "Finnish",
		},
		List{
			Id:   "fr",
			Name: "French",
		},
		List{
			Id:   "de",
			Name: "Germany",
		},
		List{
			Id:   "el",
			Name: "Greek",
		},
		List{
			Id:   "hu",
			Name: "Hungarian",
		},
		List{
			Id:   "id",
			Name: "Indonesian",
		},
		List{
			Id:   "it",
			Name: "Italian",
		},
		List{
			Id:   "ja",
			Name: "Japanese",
		},
		List{
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "lv",
			Name: "Latvian",
		},
		List{
			Id:   "lt",
			Name: "Lithuanian",
		},
		List{
			Id:   "nb",
			Name: "Norwegian",
		},
		List{
			Id:   "pl",
			Name: "Polish",
		},
		List{
			Id:   "pt",
			Name: "Portugese",
		},
		List{
			Id:   "ro",
			Name: "Romanian",
		},
		List{
			Id:   "ru",
			Name: "Russian",
		},
		List{
			Id:   "sk",
			Name: "Slovak",
		},
		List{
			Id:   "sl",
			Name: "Slovenian",
		},
		List{
			Id:   "es",
			Name: "Spanish",
		},
		List{
			Id:   "sv",
			Name: "Swedish",
		},
		List{
			Id:   "tr",
			Name: "Turkish",
		},
		List{
			Id:   "uk",
			Name: "Ukrainian",
		},
	}
	if listType == "sl" {
		auto := List{
			Id:   "auto",
			Name: "Detect Language",
		}
		ListData = append(ListData, auto)
	}
	return ListData
}
