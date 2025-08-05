package handler

import "github.com/gofiber/fiber/v2"

func InitHandlers(app *fiber.App, messageHandler *MessageHandler, schedulerHandler *SchedulerHandler) {
	scheduler := app.Group("/scheduler")
	scheduler.Post("/start", schedulerHandler.Start)
	scheduler.Post("/stop", schedulerHandler.Stop)

	app.Get("/messages", messageHandler.GetMessages)
}
