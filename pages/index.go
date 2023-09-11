package pages

import (
	"codeberg.org/aryak/libmozhi"
	"github.com/gofiber/fiber/v2"
	"os"
)

func envTrueNoExist(env string) bool {
	if _, ok := os.LookupEnv(env); ok == false || os.Getenv(env) == "true" {
		return true
	}
	return false
}

func engineList() map[string]string {
	engines := map[string]string{"google":"Google", "deepl": "DeepL", "duckduckgo": "DuckDuckGo", "libre": "LibreTranslate", "mymemory": "MyMemory", "reverso": "Reverso", "watson": "Watson", "yandex": "Yandex"}
	if envTrueNoExist("MOZHI_GOOGLE_ENABLED") == false {
		delete(engines,"google")
	} else if envTrueNoExist("MOZHI_DEEPL_ENABLED") == false {
		delete(engines,"deepl")
	} else if envTrueNoExist("MOZHI_DUCKDUCKGO_ENABLED") == false {
		delete(engines,"duckduckgo")
	} else if envTrueNoExist("MOZHI_LIBRETRANSLATE_ENABLED") == false || envTrueNoExist("MOZHI_LIBRETRANSLATE_URL") {
		delete(engines,"libre")
	} else if envTrueNoExist("MOZHI_MYMEMORY_ENABLED") == false {
		delete(engines,"mymemory")
	} else if envTrueNoExist("MOZHI_REVERSO_ENABLED") == false {
		delete(engines,"reverso")
	} else if envTrueNoExist("MOZHI_WATSON_ENABLED") == false {
		delete(engines,"watson")
	} else if envTrueNoExist("MOZHI_YANDEX_ENABLED") == false {
		delete(engines,"yandex")
	}
	return engines
} 

func HandleIndex(c *fiber.Ctx) error {
	engines := engineList()
	var engine string
	var originalText string
	if c.Query("engine") == "" {
		engine = "google"
	}
	if c.Query("engine") != "" {
		for key, _ := range engines {
			if c.Query("engine") == key {
				engine = c.Query("engine")
			}
		}
		if engine == "" {
			engine = "google"
		}
	}

	sourceLanguages, err1 := libmozhi.LangList(engine, "sl")
	targetLanguages, err2 := libmozhi.LangList(engine, "tl")
	if err1 != nil || err2 != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err1.Error()+err2.Error())
	}

	originalText = c.Query("text")
	to := c.Query("to")
	from := c.Query("from")
	var translation libmozhi.LangOut
	var tlerr error
	var ttsFrom string
	var ttsTo string
	if engine != "" && originalText != "" && from != "" && to != "" {
		translation, tlerr = libmozhi.Translate(engine, to, from, originalText)
		if tlerr != nil {
			return fiber.NewError(fiber.StatusInternalServerError, tlerr.Error())
		}
		if engine == "google" || engine == "reverso" {
			if from == "auto" {
				//ttsFrom = "/api/tts?lang="+translation.AutoDetect+"&engine="+engine+"&text="+originalText
				ttsFrom = ""
			} else {
				ttsFrom = "/api/tts?lang=" + from + "&engine=" + engine + "&text=" + originalText
			}
			ttsTo = "/api/tts?lang=" + to + "&engine=" + engine + "&text=" + translation.OutputText
		}
	}
	return c.Render("index", fiber.Map{
		"Engine":          engine,
		"enginesNames":    engines,
		"SourceLanguages": sourceLanguages,
		"TargetLanguages": targetLanguages,
		"OriginalText":    originalText,
		"Translation":     translation,
		"From":            from,
		"To":              to,
		"TtsFrom":         ttsFrom,
		"TtsTo":           ttsTo,
		"defaultLang":     "en",
	})
}
