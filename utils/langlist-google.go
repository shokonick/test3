package utils

func LangListGoogle(listType string) []List {
	// IDs got from original simplytranslate-web and trial and error. Usually first three letters of language.
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
			Id:   "ay",
			Name: "Aymara",
		},
		List{
			Id:   "az",
			Name: "Azerbaijani",
		},
		List{
			Id:   "bm",
			Name: "Bambara",
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
			Id:   "bho",
			Name: "Bhojpuri",
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
			Id:   "ceb",
			Name: "Cebuano",
		},
		List{
			Id:   "ny",
			Name: "Chichewa",
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
			Id:   "co",
			Name: "Corsican",
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
			Id:   "dv",
			Name: "Dhivehi",
		},
		List{
			Id:   "doi",
			Name: "Dogri",
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
			Id:   "eo",
			Name: "Esperanto",
		},
		List{
			Id:   "et",
			Name: "Estonian",
		},
		List{
			Id:   "ee",
			Name: "Ewe",
		},
		List{
			Id:   "tl",
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
			Id:   "fy",
			Name: "Frisian",
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
			Id:   "gn",
			Name: "Guarani",
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
			Id:   "haw",
			Name: "Hawaiian",
		},
		List{
			Id:   "iw",
			Name: "Hebrew",
		},
		List{
			Id:   "hi",
			Name: "Hindi",
		},
		List{
			Id:   "hmn",
			Name: "Hmong",
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
			Id:   "ilo",
			Name: "Ilocano",
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
			Id:   "jw",
			Name: "Javanese",
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
			Id:   "gom",
			Name: "Konkani",
		},
		List{
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "kri",
			Name: "Krio",
		},
		List{
			Id:   "ku",
			Name: "Kurdish (Kurmanji)",
		},
		List{
			Id:   "ckb",
			Name: "Kurdish (Sorani)",
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
			Id:   "ln",
			Name: "Lingala",
		},
		List{
			Id:   "lt",
			Name: "Lithuanian",
		},
		List{
			Id:   "lg",
			Name: "Luganda",
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
			Id:   "mi",
			Name: "Maori",
		},
		List{
			Id:   "mr",
			Name: "Marathi",
		},
		List{
			Id:   "mni-Mtei",
			Name: "Meiteilon (Manipuri)",
		},
		List{
			Id:   "lus",
			Name: "Mizo",
		},
		List{
			Id:   "mn",
			Name: "Mongolian",
		},
		List{
			Id:   "my",
			Name: "Myanmar (Burmese)",
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
			Id:   "or",
			Name: "Odia (Oriya)",
		},
		List{
			Id:   "om",
			Name: "Oromo",
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
			Name: "Portuguese",
		},
		List{
			Id:   "pa",
			Name: "Punjabi",
		},
		List{
			Id:   "qu",
			Name: "Quechua",
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
			Id:   "sm",
			Name: "Samoan",
		},
		List{
			Id:   "sa",
			Name: "Sanskrit",
		},
		List{
			Id:   "gd",
			Name: "Scots Gaelic",
		},
		List{
			Id:   "nso",
			Name: "Sepedi",
		},
		List{
			Id:   "sr",
			Name: "Serbian",
		},
		List{
			Id:   "st",
			Name: "Sesotho",
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
			Id:   "ti",
			Name: "Tigrinya",
		},
		List{
			Id:   "ts",
			Name: "Tsonga",
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
			Id:   "ak",
			Name: "Twi",
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
			Id:   "ug",
			Name: "Uyghur",
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
			Id:   "yi",
			Name: "Yiddish",
		},
		List{
			Id:   "yo",
			Name: "Yoruba",
		},
		List{
			Id:   "zu",
			Name: "Zulu",
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
