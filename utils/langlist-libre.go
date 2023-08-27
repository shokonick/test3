package utils

func LangListLibreTranslate(listType string) []List {
	ListData := []List{
		{
			Id:   "en",
			Name: "English",
		},
		{
			Id:   "ar",
			Name: "Arabic",
		},
		{
			Id:   "zh",
			Name: "Chinese",
		},
		{
			Id:   "fr",
			Name: "French",
		},
		{
			Id:   "de",
			Name: "German",
		},
		{
			Id:   "hi",
			Name: "Hindi",
		},
		{
			Id:   "id",
			Name: "Indonesian",
		},
		{
			Id:   "ga",
			Name: "Irish",
		},
		{
			Id:   "it",
			Name: "Italian",
		},
		{
			Id:   "ja",
			Name: "Japanese",
		},
		{
			Id:   "ko",
			Name: "Korean",
		},
		{
			Id:   "pl",
			Name: "Polish",
		},
		{
			Id:   "pt",
			Name: "Portuguese",
		},
		{
			Id:   "ru",
			Name: "Russian",
		},
		{
			Id:   "es",
			Name: "Spanish",
		},
		{
			Id:   "tr",
			Name: "Turkish",
		},
		{
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
