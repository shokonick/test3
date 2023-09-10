package pages

import (
	libmozhi "codeberg.org/aryak/libmozhi"
	"codeberg.org/aryak/mozhi/utils"
	"github.com/gofiber/fiber/v2"
)

// HandleSourceLanguages godoc
//
//	@Summary	Show list of available source languages for engine
//	@Param		engine	query		string	true	"Engine name"
//	@Success	200		{object}	libmozhi.List
//	@Router		/api/source_languages [get]
func HandleSourceLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine is a required query string.")
	}
	data, err := libmozhi.LangList(engine, "sl")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

// HandleTargetLanguages godoc
//
//	@Summary	Show list of available target languages for engine
//	@Param		engine	query		string	true	"Engine name"
//	@Success	200		{object}	libmozhi.List
//	@Router		/api/target_languages [get]
func HandleTargetLanguages(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	if engine == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine is a required query string.")
	}
	data, err := libmozhi.LangList(engine, "tl")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

// HandleTTS godoc
//
//	@Summary	Get Text-To-Speech for specified language using specified engine
//	@Param		engine	query	string	true	"Engine name"
//	@Param		lang	query	string	true	"Language being TTS'd"
//	@Param		text	query	string	true	"Text being TTS'd"
//	@Router		/api/tts [get]
func HandleTTS(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	lang := utils.Sanitize(c.Query("lang"), "alpha")
	text := c.Query("text")
	if engine == "" || text == "" || lang == "" {
		return fiber.NewError(fiber.StatusBadRequest, "engine, lang, text are required query strings.")
	}
	data, err := libmozhi.TTS(engine, lang, text)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "audio/mpeg")
	return c.Status(fiber.StatusOK).Send(data)
}

// HandleTranslate godoc
//
//	@Summary		Translate text
//	@Description	When engine is set to all, it will return an array of libmozhi.LangOut.
//	@Param			engine	query		string	true	"Engine name"
//	@Param			from	query		string	true	"Source language"
//	@Param			to		query		string	true	"Target language"
//	@Param			text	query		string	true	"Text being translated"
//	@Success		200		{object}	libmozhi.LangOut
//	@Router			/api/translate [get]
func HandleTranslate(c *fiber.Ctx) error {
	engine := utils.Sanitize(c.Query("engine"), "alpha")
	from := utils.Sanitize(c.Query("from"), "alpha")
	to := utils.Sanitize(c.Query("to"), "alpha")
	text := c.Query("text")
	if engine == "" || from == "" || to == "" || text == "" {
		return fiber.NewError(fiber.StatusBadRequest, "from, to, engine, text are required query strings.")
	}
	var dataarr []libmozhi.LangOut
	var data libmozhi.LangOut
	var err error
	if engine == "all" {
		dataarr = libmozhi.TranslateAll(to, from, text)
	} else {
		data, err = libmozhi.Translate(engine, to, from, text)
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
