package middleware

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func InitSwagger(app *fiber.App) {
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "../docs/swagger.json",
		Path:     "/",
	}

	app.Use(swagger.New(cfg))
}
