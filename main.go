package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/averincrud/internal/component"
	"github.com/khairulharu/averincrud/internal/config"
)

func main() {
	config := config.Get()
	component.GetDatabaseConnection(config)

	app := fiber.New()
	app.Use(logger.New())

	app.Listen(config.SRV.Host + ":" + config.SRV.Port)
}
