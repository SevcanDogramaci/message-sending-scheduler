package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/handler"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/mocks"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMessages_GivenStatus_ThenItShouldReturnMessages(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	expectedMessages := []*model.Message{{ID: "test-id", Content: "test-content", Status: model.StatusSent}}
	mockMessageService := mocks.NewMockMessageService(mockController)
	mockMessageService.
		EXPECT().
		GetMessages(model.StatusSent).
		Return(expectedMessages, nil)
	handler := handler.NewMessageHandler(mockMessageService)

	app := fiber.New()
	app.Get("/test", handler.GetMessages)

	req := httptest.NewRequest(http.MethodGet, "/test?status=SENT", nil)
	resp, _ := app.Test(req, -1)

	var actualMessages []*model.Message
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &actualMessages)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestGetMessages_GivenServiceFails_ThenItShouldReturnInternalServerError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockMessageService := mocks.NewMockMessageService(mockController)
	mockMessageService.
		EXPECT().
		GetMessages(model.StatusSent).
		Return(nil, errors.New("service error"))
	handler := handler.NewMessageHandler(mockMessageService)

	app := fiber.New()
	app.Get("/test", handler.GetMessages)

	req := httptest.NewRequest(http.MethodGet, "/test?status=SENT", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
