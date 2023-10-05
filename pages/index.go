package pages

import (
	"os"
	"slices"

	"codeberg.org/aryak/libmozhi"
	"github.com/gofiber/fiber/v2"
)

func envTrueNoExist(env string) bool {
	if _, ok := os.LookupEnv(env); ok == false || os.Getenv(env) == "true" {
		return true
	}
	return false
}

func engineList() map[string]string {
	engines := map[string]string{"all": "All Engines", "google": "Google", "deepl": "DeepL", "duckduckgo": "DuckDuckGo", "libre": "LibreTranslate", "mymemory": "MyMemory", "reverso": "Reverso", "watson": "Watson", "yandex": "Yandex"}
	if envTrueNoExist("MOZHI_GOOGLE_ENABLED") == false {
		delete(engines, "google")
	} else if envTrueNoExist("MOZHI_DEEPL_ENABLED") == false {
		delete(engines, "deepl")
	} else if envTrueNoExist("MOZHI_DUCKDUCKGO_ENABLED") == false {
		delete(engines, "duckduckgo")
	} else if envTrueNoExist("MOZHI_LIBRETRANSLATE_ENABLED") == false || envTrueNoExist("MOZHI_LIBRETRANSLATE_URL") {
		delete(engines, "libre")
	} else if envTrueNoExist("MOZHI_MYMEMORY_ENABLED") == false {
		delete(engines, "mymemory")
	} else if envTrueNoExist("MOZHI_REVERSO_ENABLED") == false {
		delete(engines, "reverso")
	} else if envTrueNoExist("MOZHI_WATSON_ENABLED") == false {
		delete(engines, "watson")
	} else if envTrueNoExist("MOZHI_YANDEX_ENABLED") == false {
		delete(engines, "yandex")
	}
	return engines
}

// DeduplicateLists deduplicates a slice of List based on the Id field
func deDuplicateLists(input []libmozhi.List) []libmozhi.List {
	// Create a map to store unique Ids
	uniqueIds := make(map[string]struct{})
	result := []libmozhi.List{}

	// Iterate over the input slice
	for _, item := range input {
		// Check if the Id is unique
		if _, found := uniqueIds[item.Id]; !found {
			// Add the Id to the map and append the List to the result slice
			uniqueIds[item.Id] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

func langListMerge(engines map[string]string) ([]libmozhi.List, []libmozhi.List) {
	sl := []libmozhi.List{}
	tl := []libmozhi.List{}
	for key, _ := range engines {
		temp, _ := libmozhi.LangList(key, "sl")
		temp2, _ := libmozhi.LangList(key, "tl")
		sl = append(sl, temp...)
		tl = append(tl, temp2...)
	}
	return deDuplicateLists(sl), deDuplicateLists(tl)
}

func getQueryOrFormValue(c *fiber.Ctx, key string) string {
	if c.Method() == "POST" {
		return c.FormValue(key)
	} else {
		return c.Query(key)
	}
}

func HandleIndex(c *fiber.Ctx) error {
	engines := engineList()
	var enginesAsArray []string
	for engine := range engines {
		enginesAsArray = append(enginesAsArray, engine)
	}

	var engine = getQueryOrFormValue(c, "engine")
	if engine == "" || !slices.Contains(enginesAsArray, engine) {
		engine = "google"
	}

	var sourceLanguages []libmozhi.List
	var targetLanguages []libmozhi.List
	if engine == "all" {
		sourceLanguages, targetLanguages = langListMerge(engines)
	} else {
		sourceLanguages, _ = libmozhi.LangList(engine, "sl")
		targetLanguages, _ = libmozhi.LangList(engine, "tl")
	}

	originalText := getQueryOrFormValue(c, "text")
	to := getQueryOrFormValue(c, "to")
	from := getQueryOrFormValue(c, "from")

	var translation libmozhi.LangOut
	var translationExists bool
	var transall []libmozhi.LangOut
	var tlerr error
	var ttsFrom string
	var ttsTo string
	if engine != "" && originalText != "" && from != "" && to != "" {
		if engine == "all" {
			transall = libmozhi.TranslateAll(to, from, originalText)
		} else {
			translation, tlerr = libmozhi.Translate(engine, to, from, originalText)
			if tlerr != nil {
				return fiber.NewError(fiber.StatusInternalServerError, tlerr.Error())
			}
			translationExists = true
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
	} else {
		translationExists = false
	}
	return c.Render("index", fiber.Map{
		"Engine":            engine,
		"enginesNames":      engines,
		"SourceLanguages":   sourceLanguages,
		"TargetLanguages":   targetLanguages,
		"OriginalText":      originalText,
		"Translation":       translation,
		"TranslationExists": translationExists,
		"TranslateAll":      transall,
		"From":              from,
		"To":                to,
		"TtsFrom":           ttsFrom,
		"TtsTo":             ttsTo,
		"defaultLang":       "en",
	})
}
