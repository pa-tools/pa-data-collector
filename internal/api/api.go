package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	app *fiber.App
)

func InitApi() {
	app = fiber.New()
	app.Use(logger.New())

	// Route Handlers
	app.Post("/post", parsePaPost)

	app.Listen(":3000")
}
