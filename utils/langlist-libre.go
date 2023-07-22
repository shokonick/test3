package utils

func LangListLibreTranslate(listType string) []List {
	var ListData = []List{
		List{
			Id:   "en",
			Name: "English",
		},
		List{
			Id:   "ar",
			Name: "Arabic",
		},
		List{
			Id:   "zh",
			Name: "Chinese",
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
			Id:   "hi",
			Name: "Hindi",
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
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "pl",
			Name: "Polish",
		},
		List{
			Id:   "pt",
			Name: "Portuguese",
		},
		List{
			Id:   "ru",
			Name: "Russian",
		},
		List{
			Id:   "es",
			Name: "Spanish",
		},
		List{
			Id:   "tr",
			Name: "Turkish",
		},
		List{
			Id:   "vi",
			Name: "Vietnamese",
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
