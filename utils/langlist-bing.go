package utils

func LangListBing(listType string) []List {
	// Converted json of https://github.com/plainheart/bing-translate-api/blob/master/src/lang.json to this format
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
			Id:   "as",
			Name: "Assamese",
		},
		List{
			Id:   "az",
			Name: "Azerbaijani",
		},
		List{
			Id:   "bn",
			Name: "Bangla",
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
			Id:   "bs",
			Name: "Bosnian",
		},
		List{
			Id:   "bg",
			Name: "Bulgarian",
		},
		List{
			Id:   "yue",
			Name: "Cantonese (Traditional)",
		},
		List{
			Id:   "ca",
			Name: "Catalan",
		},
		List{
			Id:   "lzh",
			Name: "Chinese (Literary)",
		},
		List{
			Id:   "zh-Hans",
			Name: "Chinese Simplified",
		},
		List{
			Id:   "zh-Hant",
			Name: "Chinese Traditional",
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
			Id:   "prs",
			Name: "Dari",
		},
		List{
			Id:   "dv",
			Name: "Divehi",
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
			Id:   "fo",
			Name: "Faroese",
		},
		List{
			Id:   "fj",
			Name: "Fijian",
		},
		List{
			Id:   "fil",
			Name: "Filipino",
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
			Name: "French (Canada)",
		},
		List{
			Id:   "gl",
			Name: "Galician",
		},
		List{
			Id:   "lug",
			Name: "Ganda",
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
			Name: "Haitian Creole",
		},
		List{
			Id:   "ha",
			Name: "Hausa",
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
			Id:   "mww",
			Name: "Hmong Daw",
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
			Id:   "ig",
			Name: "Igbo",
		},
		List{
			Id:   "id",
			Name: "Indonesian",
		},
		List{
			Id:   "ikt",
			Name: "Inuinnaqtun",
		},
		List{
			Id:   "iu",
			Name: "Inuktitut",
		},
		List{
			Id:   "iu-Latn",
			Name: "Inuktitut (Latin)",
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
			Id:   "kk",
			Name: "Kazakh",
		},
		List{
			Id:   "km",
			Name: "Khmer",
		},
		List{
			Id:   "rw",
			Name: "Kinyarwanda",
		},
		List{
			Id:   "tlh-Latn",
			Name: "Klingon (Latin)",
		},
		List{
			Id:   "gom",
			Name: "Konkani",
		},
		List{
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "ku",
			Name: "Kurdish (Central)",
		},
		List{
			Id:   "kmr",
			Name: "Kurdish (Northern)",
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
			Id:   "lv",
			Name: "Latvian",
		},
		List{
			Id:   "ln",
			Name: "Lingala",
		},
		List{
			Id:   "lt",
			Name: "Lithuanian",
		},
		List{
			Id:   "dsb",
			Name: "Lower Sorbian",
		},
		List{
			Id:   "mk",
			Name: "Macedonian",
		},
		List{
			Id:   "mai",
			Name: "Maithili",
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
			Id:   "mr",
			Name: "Marathi",
		},
		List{
			Id:   "mn-Cyrl",
			Name: "Mongolian (Cyrillic)",
		},
		List{
			Id:   "mn-Mong",
			Name: "Mongolian (Traditional)",
		},
		List{
			Id:   "my",
			Name: "Myanmar (Burmese)",
		},
		List{
			Id:   "mi",
			Name: "Māori",
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
			Id:   "nya",
			Name: "Nyanja",
		},
		List{
			Id:   "or",
			Name: "Odia",
		},
		List{
			Id:   "ps",
			Name: "Pashto",
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
			Id:   "pt",
			Name: "Portuguese (Brazil)",
		},
		List{
			Id:   "pt-PT",
			Name: "Portuguese (Portugal)",
		},
		List{
			Id:   "pa",
			Name: "Punjabi",
		},
		List{
			Id:   "otq",
			Name: "Querétaro Otomi",
		},
		List{
			Id:   "ro",
			Name: "Romanian",
		},
		List{
			Id:   "run",
			Name: "Rundi",
		},
		List{
			Id:   "ru",
			Name: "Russian",
		},
		List{
			Id:   "sm",
			Name: "Samoan",
		},
		List{
			Id:   "sr-Cyrl",
			Name: "Serbian (Cyrillic)",
		},
		List{
			Id:   "sr-Latn",
			Name: "Serbian (Latin)",
		},
		List{
			Id:   "st",
			Name: "Sesotho",
		},
		List{
			Id:   "nso",
			Name: "Sesotho sa Leboa",
		},
		List{
			Id:   "tn",
			Name: "Setswana",
		},
		List{
			Id:   "sn",
			Name: "Shona",
		},
		List{
			Id:   "sd",
			Name: "Sindhi",
		},
		List{
			Id:   "si",
			Name: "Sinhala",
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
			Id:   "so",
			Name: "Somali",
		},
		List{
			Id:   "es",
			Name: "Spanish",
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
			Id:   "ty",
			Name: "Tahitian",
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
			Id:   "bo",
			Name: "Tibetan",
		},
		List{
			Id:   "ti",
			Name: "Tigrinya",
		},
		List{
			Id:   "to",
			Name: "Tongan",
		},
		List{
			Id:   "tr",
			Name: "Turkish",
		},
		List{
			Id:   "tk",
			Name: "Turkmen",
		},
		List{
			Id:   "uk",
			Name: "Ukrainian",
		},
		List{
			Id:   "hsb",
			Name: "Upper Sorbian",
		},
		List{
			Id:   "ur",
			Name: "Urdu",
		},
		List{
			Id:   "ug",
			Name: "Uyghur",
		},
		List{
			Id:   "uz",
			Name: "Uzbek (Latin)",
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
			Id:   "yo",
			Name: "Yoruba",
		},
		List{
			Id:   "yua",
			Name: "Yucatec Maya",
		},
		List{
			Id:   "zu",
			Name: "Zulu",
		},
	}
	return ListData
}
