package serve

import (
	"html/template"
	"log"
	"os"
	"runtime"

	"codeberg.org/aryak/mozhi/pages"
	"codeberg.org/aryak/mozhi/utils"
	_ "codeberg.org/aryak/mozhi/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	//"github.com/gofiber/fiber/v2/middleware/limiter"
	//	For debugging purposes
	//	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"

	_ "github.com/joho/godotenv/autoload"
)

//	@title			Mozhi API
//	@version		1.0
//	@description	API for Mozhi, the alternative-frontend for many translation engines.
//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.txt
//	@BasePath		/api
func Serve(port string) {
	engine := html.New("./views", ".html")

	engine.AddFunc(
		// Add unescape function. This is needed to render HTML from Markdown.
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: false,
		AppName: "Mozhi",
		// kind of screwed up way to fix rate limits
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	// For debugging purposes
	// app.Use(logger.New(logger.Config{
	// 	Format: "[${ip}]:${port} ${status} - ${method} ${path} ${queryParams}\n",
	// }))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(recover.New())

	//ratelimiter := limiter.New(limiter.Config{
	//	Max:        5,
	//	Expiration: 5 * time.Minute,
	//	LimitReached: func(c *fiber.Ctx) error {
	//		return c.Status(429).Render("ratelimit_gt", fiber.Map{
	//			"Title":  "Rate limit exceeded",
	//		})
	//	},
	//})

	staticConfig := fiber.Static{
		Compress: true,
		// Cache-Control: max-age=31536000
		MaxAge: 31536000,
	}

	// add global headers
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("Referrer-Policy", "no-referrer")
		c.Set("Content-Security-Policy", "default-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; script-src 'self'; frame-ancestors 'self'; form-action 'self'; base-uri 'self'; connect-src 'self';")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		return c.Next()
	})

	app.Get("/", pages.HandleIndex)
	app.Static("/css", "./public/css", staticConfig)
	app.Static("/robots.txt", "./public/robots.txt", staticConfig)
	app.Static("/favicon.ico", "./public/assets/favicon.ico", staticConfig)
	app.Static("/logo.svg", "./public/assets/logo.svg", staticConfig)
	// app.Get("/about", pages.HandleAbout)

	api := app.Group("/api")
	api.Get("/translate", pages.HandleTranslate)
	api.Get("/source_languages", pages.HandleSourceLanguages)
	api.Get("/target_languages", pages.HandleTargetLanguages)
	api.Get("/tts", pages.HandleTTS)
	api.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"version":      utils.Version(),
			"fiberversion": fiber.Version,
			"goversion":    runtime.Version(),
		})
	})
	api.Get("/swagger/*", swagger.HandlerDefault) // default

	val, ok := os.LookupEnv("MOZHI_PORT")
	if !ok {
		val = "3000"
	}
	if port != "" {
		val = port
	}
	log.Fatal(app.Listen(":" + val))
}
