package utils

type List struct {
	Name string
	Id   string
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
