package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/averincrud/internal/api"
	"github.com/khairulharu/averincrud/internal/component"
	"github.com/khairulharu/averincrud/internal/config"
	"github.com/khairulharu/averincrud/internal/repository"
	"github.com/khairulharu/averincrud/internal/service"
)

func main() {
	config := config.Get()
	dbConnection := component.GetDatabaseConnection(config)

	patientRepository := repository.NewPatient(dbConnection)

	patientService := service.NewPatient(patientRepository)

	app := fiber.New()
	app.Use(logger.New())

	api.NewPatient(patientService, app)

	app.Listen(config.SRV.Host + ":" + config.SRV.Port)
}
