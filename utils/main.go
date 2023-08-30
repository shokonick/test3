package utils

import (
	"errors"
	"os"
)

type List struct {
	Name string
	Id   string
}
type LangOut struct {
	Engine     string `json:"engine"`
	AutoDetect string `json:"detected"`
	OutputText string `json:"translated-text"`
	SourceLang string `json:"source_language"`
	TargetLang string `json:"target_language"`
}

func EnvTrueNoExist(env string) bool {
	if _, ok := os.LookupEnv(env); ok == false || os.Getenv(env) == "true" {
		return true
	}
	return false
}

func LangList(engine string, listType string) ([]List, error) {
	var data []List
	if listType != "sl" && listType != "tl" {
		return []List{}, errors.New("list type invalid: either give tl for target languages or sl for source languages.")
	}
	if engine == "google" && EnvTrueNoExist("MOZHI_GOOGLE_ENABLED") {
		data = LangListGoogle(listType)
	} else if engine == "libre" && EnvTrueNoExist("MOZHI_LIBRETRANSLATE_ENABLED") {
		if EnvTrueNoExist("MOZHI_LIBRETRANSLATE_URL") {
			return []List{}, errors.New("Please set MOZHI_LIBRETRANSLATE_URL if you want to use libretranslate. Example: MOZHI_LIBRETRANSLATE_URL=https://lt.psf.lt")
		}
		data = LangListLibreTranslate(listType)
	} else if engine == "reverso" && EnvTrueNoExist("MOZHI_REVERSO_ENABLED") {
		data = LangListReverso(listType)
	} else if engine == "deepl" && EnvTrueNoExist("MOZHI_DEEPL_ENABLED") {
		data = LangListDeepl(listType)
	} else if engine == "watson" && EnvTrueNoExist("MOZHI_WATSON_ENABLED") {
		data = LangListWatson(listType)
	} else if engine == "yandex" && EnvTrueNoExist("MOZHI_YANDEX_ENABLED") {
		data = LangListYandex(listType)
	} else if engine == "mymemory" && EnvTrueNoExist("MOZHI_MYMEMORY_ENABLED") {
		data = LangListMyMemory(listType)
	} else if engine == "duckduckgo" && EnvTrueNoExist("MOZHI_DUCKDUCKGO_ENABLED") {
		data = LangListDuckDuckGo(listType)
	} else {
		return []List{}, errors.New("Engine does not exist or has been disabled.")
	}
	return data, nil
}

// General function to translate stuff so there is no need for a huge if-block everywhere
func Translate(engine string, to string, from string, text string) (LangOut, error) {
	var err error
	var data LangOut
	if engine == "google" && EnvTrueNoExist("MOZHI_GOOGLE_ENABLED") {
		data, err = TranslateGoogle(to, from, text)
	} else if engine == "libre" && EnvTrueNoExist("MOZHI_LIBRETRANSLATE_ENABLED") {
		if os.Getenv("MOZHI_LIBRETRANSLATE_URL") == "" {
			return LangOut{}, errors.New("Please set MOZHI_LIBRETRANSLATE_URL if you want to use libretranslate. Example: MOZHI_LIBRETRANSLATE_URL=https://lt.psf.lt")
		}
		data, err = TranslateLibreTranslate(to, from, text)
	} else if engine == "reverso" && EnvTrueNoExist("MOZHI_REVERSO_ENABLED") {
		data, err = TranslateReverso(to, from, text)
	} else if engine == "deepl" && EnvTrueNoExist("MOZHI_DEEPL_ENABLED") {
		data, err = TranslateDeepl(to, from, text)
	} else if engine == "watson" && EnvTrueNoExist("MOZHI_WATSON_ENABLED") {
		data, err = TranslateWatson(to, from, text)
	} else if engine == "yandex" && EnvTrueNoExist("MOZHI_YANDEX_ENABLED") {
		data, err = TranslateYandex(to, from, text)
	} else if engine == "mymemory" && EnvTrueNoExist("MOZHI_MYMEMORY_ENABLED") {
		data, err = TranslateMyMemory(to, from, text)
	} else if engine == "duckduckgo" && EnvTrueNoExist("MOZHI_DUCKDUCKGO_ENABLED") {
		data, err = TranslateDuckDuckGo(to, from, text)
	} else {
		return LangOut{}, errors.New("Engine does not exist or has been disabled.")
	}
	if err != nil {
		return LangOut{}, err
	}
	return data, nil
}

func TTS(engine string, lang string, text string) ([]byte, error) {
	var err error
	var data []byte
	if engine == "google" && EnvTrueNoExist("MOZHI_GOOGLE_ENABLED") {
		data, err = TTSGoogle(lang, text)
	} else if engine == "reverso" && EnvTrueNoExist("MOZHI_REVERSO_ENABLED") {
		data, err = TTSReverso(lang, text)
	} else {
		return []byte(""), errors.New("Engine does not exist and/or doesn't support TTS and/or has been disabled.")
	}
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
