package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

var (
	app *fiber.App
)

func InitApi() {
	app = fiber.New()
	app.Use(logger.New())

	// Route Handlers
	app.Post("/post", parsePaPost)
	app.Get("/devices", getDevices)

	viper.SetDefault("port", 3000)
	viper.BindEnv("port")
	app.Listen(":" + viper.GetString("port"))
}
