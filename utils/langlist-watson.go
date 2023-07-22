package utils

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
