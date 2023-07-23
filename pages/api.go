package pages

import (
	"codeberg.org/aryak/simplytranslate/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleSourceLanguages(c *fiber.Ctx) error {
	engine := c.Query("engine")
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
	engine := c.Query("engine")
	if engine == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var data []utils.List
	if engine == "google" {
		data = utils.LangListGoogle("tl")
	} else if engine == "libre" {
		data = utils.LangListLibreTranslate("sl")
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
	engine := c.Query("engine")
	lang := c.Query("lang")
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
