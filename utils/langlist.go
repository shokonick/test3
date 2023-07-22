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
func LangListWatson(listType string) []List {
	// IDs got from https://www.loc.gov/standards/iso639-2/php/code_list.php and tested to make sure they work. Exceptions fr-CA zh-CN/TW
	var ListData = []List{
		List{
			Id:   "ar",
			Name: "Arabic",
		},
		List{
			Id:   "ba",
			Name: "Basque",
		},
		List{
			Id:   "bn",
			Name: "Bengali",
		},
		List{
			Id:   "bs",
			Name: "Bosnian",
		},
		List{
			Id:   "bg",
			Name: "Bulgarian",
		},
		List{
			Id:   "ca",
			Name: "Catalan",
		},
		List{
			Id:   "zh",
			Name: "Chinese (Simplified)",
		},
		List{
			Id:   "zh-TW",
			Name: "Chinese (Traditional)",
		},
		List{
			Id:   "hr",
			Name: "Croatian",
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
			Id:   "fr-CA",
			Name: "Canadian French",
		},
		List{
			Id:   "de",
			Name: "German",
		},
		List{
			Id:   "el",
			Name: "Greek",
		},
		List{
			Id:   "gu",
			Name: "Gujarati",
		},
		List{
			Id:   "he",
			Name: "Hebrew",
		},
		List{
			Id:   "hi",
			Name: "Hindi",
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
			Id:   "ga",
			Name: "Irish",
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
			Id:   "kn",
			Name: "Kannada",
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
			Id:   "ms",
			Name: "Malay",
		},
		List{
			Id:   "ml",
			Name: "Malayalam",
		},
		List{
			Id:   "mt",
			Name: "Maltese",
		},
		List{
			Id:   "mr",
			Name: "Marathi",
		},
		List{
			Id:   "cnr",
			Name: "Montenegrin",
		},
		List{
			Id:   "ne",
			Name: "Nepali",
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
			Id:   "pa",
			Name: "Punjabi",
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
			Id:   "sr",
			Name: "Serbian",
		},
		List{
			Id:   "si",
			Name: "Sinhalese",
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
			Id:   "ta",
			Name: "Tamil",
		},
		List{
			Id:   "te",
			Name: "Telugu",
		},
		List{
			Id:   "th",
			Name: "Thai",
		},
		List{
			Id:   "tr",
			Name: "Turkish",
		},
		List{
			Id:   "uk",
			Name: "Ukrainian",
		},
		List{
			Id:   "ur",
			Name: "Urdu",
		},
		List{
			Id:   "vi",
			Name: "Vietnamese",
		},
		List{
			Id:   "cy",
			Name: "Welsh",
		},
	}
	return ListData
}
func LangListReverso(listType string) []List {
	// IDs got from original simplytranslate-web and trial and error. Usually first three letters of language.
	var ListData = []List{
		List{
			Id:   "ar",
			Name: "Arabic",
		},
		List{
			Id:   "zh",
			Name: "Chinese (Simplified)",
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
			Id:   "fr",
			Name: "French",
		},
		List{
			Id:   "de",
			Name: "German",
		},
		List{
			Id:   "el",
			Name: "Greek",
		},
		List{
			Id:   "he",
			Name: "Hebrew",
		},
		List{
			Id:   "hi",
			Name: "Hindi",
		},
		List{
			Id:   "hu",
			Name: "Hungarian",
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
			Id:   "per",
			Name: "Persian",
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
			Id:   "es",
			Name: "Spanish",
		},
		List{
			Id:   "sv",
			Name: "Swedish",
		},
		List{
			Id:   "th",
			Name: "Thai",
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
