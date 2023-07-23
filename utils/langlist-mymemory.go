package utils

func LangListMyMemory(listType string) []List {
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
			Id:   "bjs",
			Name: "Bajan",
		},
		List{
			Id:   "rm",
			Name: "Balkan Gipsy",
		},
		List{
			Id:   "eu",
			Name: "Basque",
		},
		List{
			Id:   "bem",
			Name: "Bemba",
		},
		List{
			Id:   "bn",
			Name: "Bengali",
		},
		List{
			Id:   "be",
			Name: "Bielarus",
		},
		List{
			Id:   "bi",
			Name: "Bislama",
		},
		List{
			Id:   "bs",
			Name: "Bosnian",
		},
		List{
			Id:   "br",
			Name: "Breton",
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
			Id:   "ch",
			Name: "Chamorro",
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
			Id:   "zdj",
			Name: "Comorian",
		},
		List{
			Id:   "cop",
			Name: "Coptic",
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
			Id:   "dz",
			Name: "Dzongkha",
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
			Id:   "fnG",
			Name: "Fanagalo",
		},
		List{
			Id:   "fo",
			Name: "Faroese",
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
			Id:   "grc",
			Name: "Greek (Classical)",
		},
		List{
			Id:   "gu",
			Name: "Gujarati",
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
			Id:   "is",
			Name: "Icelandic",
		},
		List{
			Id:   "id",
			Name: "Indonesian",
		},
		List{
			Id:   "kl",
			Name: "Inuktitut (Greenland)",
		},
		List{
			Id:   "ga",
			Name: "Irish Gaelic",
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
			Id:   "kea",
			Name: "Kabuverdianu",
		},
		List{
			Id:   "kab",
			Name: "Kabylian",
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
			Id:   "rn",
			Name: "Kirundi",
		},
		List{
			Id:   "ko",
			Name: "Korean",
		},
		List{
			Id:   "ku",
			Name: "Kurdish",
		},
		List{
			Id:   "ckb",
			Name: "Kurdish Sorani",
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
			Id:   "dv",
			Name: "Maldivian",
		},
		List{
			Id:   "mt",
			Name: "Maltese",
		},
		List{
			Id:   "gv",
			Name: "Manx Gaelic",
		},
		List{
			Id:   "mi",
			Name: "Maori",
		},
		List{
			Id:   "mh",
			Name: "Marshallese",
		},
		List{
			Id:   "men",
			Name: "Mende",
		},
		List{
			Id:   "mn",
			Name: "Mongolian",
		},
		List{
			Id:   "mfe",
			Name: "Morisyen",
		},
		List{
			Id:   "ne",
			Name: "Nepali",
		},
		List{
			Id:   "niu",
			Name: "Niuean",
		},
		List{
			Id:   "no",
			Name: "Norwegian",
		},
		List{
			Id:   "ny",
			Name: "Nyanja",
		},
		List{
			Id:   "ur",
			Name: "Pakistani",
		},
		List{
			Id:   "pau",
			Name: "Palauan",
		},
		List{
			Id:   "pa",
			Name: "Panjabi",
		},
		List{
			Id:   "pap",
			Name: "Papiamentu",
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
			Id:   "pis",
			Name: "Pijin",
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
			Id:   "pot",
			Name: "Potawatomi",
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
			Id:   "sg",
			Name: "Sango",
		},
		List{
			Id:   "gd",
			Name: "Scots Gaelic",
		},
		List{
			Id:   "sr",
			Name: "Serbian",
		},
		List{
			Id:   "sn",
			Name: "Shona",
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
			Id:   "st",
			Name: "Sotho Southern",
		},
		List{
			Id:   "es",
			Name: "Spanish",
		},
		List{
			Id:   "srn",
			Name: "Sranan Tongo",
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
			Id:   "syc",
			Name: "Syriac (Aramic)",
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
			Id:   "tmh",
			Name: "Tamashek (Tuareg)",
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
			Id:   "tet",
			Name: "Tetum",
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
			Id:   "tpi",
			Name: "Tok Pisin",
		},
		List{
			Id:   "tkl",
			Name: "Tokelauan",
		},
		List{
			Id:   "to",
			Name: "Tongan",
		},
		List{
			Id:   "tn",
			Name: "Tswana",
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
			Id:   "tvl",
			Name: "Tuvaluan",
		},
		List{
			Id:   "uk",
			Name: "Ukrainian",
		},
		List{
			Id:   "ppk",
			Name: "Uma",
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
			Id:   "wls",
			Name: "Wallisian",
		},
		List{
			Id:   "cy",
			Name: "Welsh",
		},
		List{
			Id:   "wo",
			Name: "Wolof",
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
			Id:   "zu",
			Name: "Zulu",
		},
	}
	return ListData
}
