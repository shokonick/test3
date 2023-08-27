package pages

import (
	"codeberg.org/aryak/mozhi/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleSourceLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine is a required query string.")
	}
	data, err := utils.LangList(engine, "sl")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func HandleTargetLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine is a required query string.")
	}
	data, err := utils.LangList(engine, "tl")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func HandleTTS(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	lang := utils.Sanitize(c.Query("lang"), "alpha")
	text := c.Query("text")
	if engine == "" || text == "" || lang == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine, lang, text are required query strings.")
	}
	data, err := utils.TTS(engine, lang, text)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "audio/mpeg")
	return c.Send(data)
}

func HandleTranslate(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	from := utils.Sanitize(c.Query("from"), "alpha")
	to := utils.Sanitize(c.Query("to"), "alpha")
	text := c.Query("text")
	if engine == "" || from == "" || to == "" || text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "from, to, engine, text are required query strings.")
	}
	var dataarr []utils.LangOut
	var data utils.LangOut
	var err error
	if engine == "all" {
		dataarr = utils.TranslateAll(to, from, text)
	} else {
		data, err = utils.Translate(engine, to, from, text)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}
	if engine == "all" {
		return c.JSON(dataarr)
	} else {
		return c.JSON(data)
	}
}
