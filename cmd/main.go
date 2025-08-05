package main

import (
	"log"
	"os"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	client "github.com/SevcanDogramaci/message-sending-scheduler/internal/client/webhook-site"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/handler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/repository"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/scheduler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/service"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/couchbase"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env := os.Getenv("APP_ENV")
	appConfig, err := config.InitConfigs(env)
	if err != nil {
		log.Fatalf("Failed to initialize configurations: %v", err)
	}

	cb, err := couchbase.NewCouchbase(appConfig.Couchbase)
	if err != nil {
		log.Fatalf("Failed to connect to Couchbase: %v", err)
	}

	webhookClient := client.NewWebhookSiteClient(appConfig.Webhook)
	messageRepository := repository.NewMessageRepository(cb.Cluster)

	messageService := service.NewMessageService(messageRepository, webhookClient)
	schedulerService := scheduler.NewScheduler(appConfig.Scheduler, messageService)

	messageHandler := handler.NewMessageHandler(messageService)
	schedulerHandler := handler.NewSchedulerHandler(schedulerService)

	app := fiber.New()
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "../docs/swagger.json",
		Path:     "/",
	}

	app.Use(swagger.New(cfg))

	handler.InitHandlers(app, messageHandler, schedulerHandler)
	schedulerService.Start()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server stopped... Err: %v", err)
	}
}
