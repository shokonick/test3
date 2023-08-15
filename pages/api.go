package pages

import (
	"codeberg.org/aryak/simplytranslate/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleSourceLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var data []utils.List
	if engine == "google" {
		data = utils.LangListGoogle("sl")
	} else if engine == "libre" {
		data = utils.LangListLibreTranslate("sl")
	} else if engine == "reverso" {
		data = utils.LangListReverso("sl")
	} else if engine == "deepl" {
		data = utils.LangListDeepl("sl")
	} else if engine == "watson" {
		data = utils.LangListWatson("sl")
	} else if engine == "yandex" {
		data = utils.LangListYandex("sl")
	} else if engine == "mymemory" {
		data = utils.LangListMyMemory("sl")
	}
	return c.JSON(data)
}
func HandleTargetLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var data []utils.List
	if engine == "google" {
		data = utils.LangListGoogle("tl")
	} else if engine == "libre" {
		data = utils.LangListLibreTranslate("tl")
	} else if engine == "reverso" {
		data = utils.LangListReverso("tl")
	} else if engine == "deepl" {
		data = utils.LangListDeepl("tl")
	} else if engine == "watson" {
		data = utils.LangListWatson("tl")
	} else if engine == "yandex" {
		data = utils.LangListYandex("tl")
	} else if engine == "mymemory" {
		data = utils.LangListMyMemory("tl")
	}
	return c.JSON(data)
}
func HandleTTS(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	lang := utils.Sanitize(c.Query("lang"), "alpha")
	text := c.Query("text")
	// Why does go not have an andor statement :(
	if engine == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	} else if text == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	} else if lang == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var data string
	if engine == "google" {
		data = utils.TTSGoogle(lang, text)
	} else if engine == "reverso" {
		data = utils.TTSReverso(lang, text)
	}
	c.Set("Content-Type", "audio/mpeg")
	return c.Send([]byte(data))
}
func HandleTranslate(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	from := utils.Sanitize(c.Query("from"), "alpha")
	to := utils.Sanitize(c.Query("to"), "alpha")
	text := c.Query("text")
	if engine == "" && from == "" && to == "" && text == ""{
		return fiber.NewError(fiber.StatusBadRequest, "from, to, engine, text are required query strings.")
	}
	var data utils.LangOut
	var err error
	if engine == "google" {
		data, err = utils.TranslateGoogle(to, from, text)
	} else if engine == "libre" {
		data, err = utils.TranslateLibreTranslate(to, from, text)
	} else if engine == "reverso" {
		data, err = utils.TranslateReverso(to, from, text)
	} else if engine == "deepl" {
		data, err = utils.TranslateDeepl(to, from, text)
	} else if engine == "watson" {
		data, err = utils.TranslateWatson(to, from, text)
	} else if engine == "yandex" {
		data, err = utils.TranslateYandex(to, from, text)
	} else if engine == "mymemory" {
		data, err = utils.TranslateMyMemory(to, from, text)
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data.SourceLang = from
	data.TargetLang = to
	return c.JSON(data)
}
