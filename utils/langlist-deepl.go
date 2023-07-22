package utils

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
