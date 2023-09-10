package pages

import (
	"github.com/gofiber/fiber/v2"
	"codeberg.org/aryak/libmozhi"
)

func HandleIndex(c *fiber.Ctx) error {
	engines := []string{"google", "deepl", "duckduckgo", "libretranslate", "mymemory", "reverso", "watson", "yandex"}
	var engine string
	var originalText string
	if c.Query("engine") == "" {
		engine = "google"
	}
	if c.Query("engine") != "" {
		for _, name := range engines {
			if c.Query("engine") == name {
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
				ttsFrom = "/api/tts?lang="+from+"&engine="+engine+"&text="+originalText
			}
			ttsTo = "/api/tts?lang="+to+"&engine="+engine+"&text="+translation.OutputText
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
	})
}
