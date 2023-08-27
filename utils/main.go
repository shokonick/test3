package utils

import (
	"errors"
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

func LangList(engine string, listType string) ([]List, error) {
	var data []List
	if listType != "sl" && listType != "tl" {
		return []List{}, errors.New("list type invalid: either give tl for target languages or sl for source languages.")
	}
	if engine == "google" {
		data = LangListGoogle("sl")
	} else if engine == "libre" {
		data = LangListLibreTranslate("sl")
	} else if engine == "reverso" {
		data = LangListReverso("sl")
	} else if engine == "deepl" {
		data = LangListDeepl("sl")
	} else if engine == "watson" {
		data = LangListWatson("sl")
	} else if engine == "yandex" {
		data = LangListYandex("sl")
	} else if engine == "mymemory" {
		data = LangListMyMemory("sl")
	} else if engine == "duckduckgo" {
		data = LangListDuckDuckGo("sl")
	} else {
		return []List{}, errors.New("Engine does not exist.")
	}
	return data, nil
}

// General function to translate stuff so there is no need for a huge if-block everywhere
func Translate(engine string, to string, from string, text string) (LangOut, error) {
	var err error
	var data LangOut
	if engine == "google" {
		data, err = TranslateGoogle(to, from, text)
	} else if engine == "libre" {
		data, err = TranslateLibreTranslate(to, from, text)
	} else if engine == "reverso" {
		data, err = TranslateReverso(to, from, text)
	} else if engine == "deepl" {
		data, err = TranslateDeepl(to, from, text)
	} else if engine == "watson" {
		data, err = TranslateWatson(to, from, text)
	} else if engine == "yandex" {
		data, err = TranslateYandex(to, from, text)
	} else if engine == "mymemory" {
		data, err = TranslateMyMemory(to, from, text)
	} else if engine == "duckduckgo" {
		data, err = TranslateDuckDuckGo(to, from, text)
	} else {
		return LangOut{}, errors.New("Engine does not exist.")
	}
	if err != nil {
		return LangOut{}, err
	}
	return data, nil
}

func TTS(engine string, lang string, text string) ([]byte, error) {
	var err error
	var data []byte
	if engine == "google" {
		data, err = TTSGoogle(lang, text)
	} else if engine == "reverso" {
		data, err = TTSReverso(lang, text)
	} else {
		return []byte(""), errors.New("Engine does not exist and/or doesn't support TTS.")
	}
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}
