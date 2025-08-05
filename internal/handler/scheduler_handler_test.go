package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/handler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStart_ItShouldStartScheduler(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockSchedulerService := mocks.NewMockSchedulerService(mockController)
	mockSchedulerService.
		EXPECT().
		Start().
		Return(nil)
	handler := handler.NewSchedulerHandler(mockSchedulerService)

	app := fiber.New()
	app.Post("/test", handler.Start)

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestStart_GivenServiceFails_ThenItShouldReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockSchedulerService := mocks.NewMockSchedulerService(mockController)
	mockSchedulerService.
		EXPECT().
		Start().
		Return(errors.New("service error"))
	handler := handler.NewSchedulerHandler(mockSchedulerService)

	app := fiber.New()
	app.Post("/test", handler.Start)

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestStop_ItShouldStopScheduler(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockSchedulerService := mocks.NewMockSchedulerService(mockController)
	mockSchedulerService.
		EXPECT().
		Stop().
		Return(nil)
	handler := handler.NewSchedulerHandler(mockSchedulerService)

	app := fiber.New()
	app.Post("/test", handler.Stop)

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}


func TestStop_GivenServiceFails_ThenItShouldReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockSchedulerService := mocks.NewMockSchedulerService(mockController)
	mockSchedulerService.
		EXPECT().
		Stop().
		Return(errors.New("service error"))
	handler := handler.NewSchedulerHandler(mockSchedulerService)

	app := fiber.New()
	app.Post("/test", handler.Stop)

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}