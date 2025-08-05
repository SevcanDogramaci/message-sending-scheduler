package main

import (
	"log"
	"os"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	client "github.com/SevcanDogramaci/message-sending-scheduler/internal/client/webhook_site"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/handler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/middleware"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/repository"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/scheduler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/service"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/couchbase"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/redis"
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

	redis, err := redis.NewRedis(appConfig.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	webhookClient := client.NewWebhookSiteClient(appConfig.Webhook)
	messageRepository := repository.NewMessageRepository(cb.Cluster)
	cacheRepository := repository.NewCacheRepository(redis.Rdb)

	messageService := service.NewMessageService(webhookClient, messageRepository, cacheRepository)
	schedulerService := scheduler.NewScheduler(appConfig.Scheduler, messageService)

	messageHandler := handler.NewMessageHandler(messageService)
	schedulerHandler := handler.NewSchedulerHandler(schedulerService)

	app := fiber.New(fiber.Config{ErrorHandler: middleware.InitErrorHandler})
	middleware.InitSwagger(app)

	handler.InitHandlers(app, messageHandler, schedulerHandler)
	schedulerService.Start()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server stopped... Err: %v", err)
	}
}
