package utils

import (
	"fmt"
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
func TranslateAll(to string, from string, query string) string {
	reverso := TranslateReverso(to, from, query)
	google := TranslateGoogle(to, from, query)
	fmt.Println("Google: " + google)
	fmt.Println("Reverso: " + reverso)
	return "Google: " + google + "\nReverso: " + reverso
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
