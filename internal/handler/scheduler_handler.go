package handler

import "github.com/gofiber/fiber/v2"

type SchedulerService interface {
	Start() error
	Stop() error
}

type SchedulerHandler struct {
	schedulerService SchedulerService
}

func NewSchedulerHandler(schedulerService SchedulerService) *SchedulerHandler {
	return &SchedulerHandler{schedulerService: schedulerService}
}

// Start is a function to start the message scheduler
//
//	@Summary	Start scheduler
//	@Tags		scheduler
//	@Success	200
//	@Router		/scheduler/start [post]
func (h *SchedulerHandler) Start(c *fiber.Ctx) error {
	return h.schedulerService.Start()
}

// Stop is a function to stop the message scheduler
//
//	@Summary	Stop scheduler
//	@Tags		scheduler
//	@Success	200
//	@Router		/scheduler/stop [post]
func (h *SchedulerHandler) Stop(c *fiber.Ctx) error {
	return h.schedulerService.Stop()
}
