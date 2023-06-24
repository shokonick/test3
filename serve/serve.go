package serve

import (
	"html/template"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"codeberg.org/aryak/simplytranslate/pages"
	"codeberg.org/aryak/simplytranslate/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	//"github.com/gofiber/fiber/v2/middleware/limiter"
	//	For debugging purposes
	//	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"

	_ "github.com/joho/godotenv/autoload"
)

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
		AppName: "SimplyTranslate",
		// kind of screwed up way to fix rate limits
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			link := strings.TrimPrefix(err.Error(), "Cannot GET ")
			err = ctx.Status(code).Render("error", fiber.Map{
				"error":  err,
				"link":   link,
				"branch": utils.Branch,
			})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		},
	})

	app.Use(cache.New(cache.Config{
		Expiration: 5 * time.Minute,
	}))
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
	//			"branch": utils.Branch,
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
	//app.Get("/about", pages.HandleAbout)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"version":      utils.Version(),
			"fiberversion": fiber.Version,
			"goversion":    runtime.Version(),
		})
	})

	val, ok := os.LookupEnv("SIMPLYTRANSLATE_PORT")
	if !ok {
		val = "3000"
	}
	if port != "" {
		val = port
	}
	log.Fatal(app.Listen(":" + val))
}
