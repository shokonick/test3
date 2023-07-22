package utils

func LangListYandex(listType string) []List {
	var ListData = []List{
		List{
			Id:   "af",
			Name: "Afrikaans",
		},
		List{
			Id:   "sq",
			Name: "Albanian",
		},
		List{
			Id:   "am",
			Name: "Amharic",
		},
		List{
			Id:   "ar",
			Name: "Arabic",
		},
		List{
			Id:   "hy",
			Name: "Armenian",
		},
		List{
			Id:   "az",
			Name: "Azerbaijani",
		},
		List{
			Id:   "ba",
			Name: "Bashkir",
		},
		List{
			Id:   "eu",
			Name: "Basque",
		},
		List{
			Id:   "be",
			Name: "Belarusian",
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
			Id:   "my",
			Name: "Burmese",
		},
		List{
			Id:   "ca",
			Name: "Catalan",
		},
		List{
			Id:   "ceb",
			Name: "Cebuano",
		},
		List{
			Id:   "zh",
			Name: "Chinese",
		},
		List{
			Id:   "cv",
			Name: "Chuvash",
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
			Id:   "sjn",
			Name: "Elvish (Sindarin)",
		},
		List{
			Id:   "emj",
			Name: "Emoji",
		},
		List{
			Id:   "en",
			Name: "English",
		},
		List{
			Id:   "eo",
			Name: "Esperanto",
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
			Id:   "gl",
			Name: "Galician",
		},
		List{
			Id:   "ka",
			Name: "Georgian",
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
			Id:   "ht",
			Name: "Haitian",
		},
		List{
			Id:   "he",
			Name: "Hebrew",
		},
		List{
			Id:   "mrj",
			Name: "Hill Mari",
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
			Id:   "is",
			Name: "Icelandic",
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
			Id:   "jv",
			Name: "Javanese",
		},
		List{
			Id:   "kn",
			Name: "Kannada",
		},
		List{
			Id:   "kazlat",
			Name: "Kazakh (Latin)",
		},
		List{
			Id:   "kk",
			Name: "Kazakh",
		},
		List{
			Id:   "km",
			Name: "Khmer",
		},
		List{
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "ky",
			Name: "Kyrgyz",
		},
		List{
			Id:   "lo",
			Name: "Lao",
		},
		List{
			Id:   "la",
			Name: "Latin",
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
			Id:   "lb",
			Name: "Luxembourgish",
		},
		List{
			Id:   "mk",
			Name: "Macedonian",
		},
		List{
			Id:   "mg",
			Name: "Malagasy",
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
			Id:   "mi",
			Name: "Maori",
		},
		List{
			Id:   "mr",
			Name: "Marathi",
		},
		List{
			Id:   "mhr",
			Name: "Mari",
		},
		List{
			Id:   "mn",
			Name: "Mongolian",
		},
		List{
			Id:   "ne",
			Name: "Nepali",
		},
		List{
			Id:   "no",
			Name: "Norwegian",
		},
		List{
			Id:   "pap",
			Name: "Papiamento",
		},
		List{
			Id:   "fa",
			Name: "Persian",
		},
		List{
			Id:   "pl",
			Name: "Polish",
		},
		List{
			Id:   "pt-BR",
			Name: "Portuguese (Brazilian)",
		},
		List{
			Id:   "pt",
			Name: "Portuguese",
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
			Id:   "gd",
			Name: "Scottish Gaelic",
		},
		List{
			Id:   "sr-Latn",
			Name: "Serbian (Latin)",
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
			Id:   "su",
			Name: "Sundanese",
		},
		List{
			Id:   "sw",
			Name: "Swahili",
		},
		List{
			Id:   "sv",
			Name: "Swedish",
		},
		List{
			Id:   "tl",
			Name: "Tagalog",
		},
		List{
			Id:   "tg",
			Name: "Tajik",
		},
		List{
			Id:   "ta",
			Name: "Tamil",
		},
		List{
			Id:   "tt",
			Name: "Tatar",
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
			Id:   "udm",
			Name: "Udmurt",
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
			Id:   "uzbcyr",
			Name: "Uzbek (Cyrillic)",
		},
		List{
			Id:   "uz",
			Name: "Uzbek",
		},
		List{
			Id:   "vi",
			Name: "Vietnamese",
		},
		List{
			Id:   "cy",
			Name: "Welsh",
		},
		List{
			Id:   "xh",
			Name: "Xhosa",
		},
		List{
			Id:   "sah",
			Name: "Yakut",
		},
		List{
			Id:   "yi",
			Name: "Yiddish",
		},
		List{
			Id:   "zu",
			Name: "Zulu",
		},
	}
	return ListData
}
