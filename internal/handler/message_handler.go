package handler

import (
	"errors"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2"
)

type MessageService interface {
	GetMessages(status model.Status) ([]model.Message, error)
}

type MessageHandler struct {
	messageService MessageService
}

func NewMessageHandler(messageService MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

// GetMessages is a function to get messages from database
//
//	@Summary	Get messages by status
//	@Tags		messages
//	@Produce	json
//	@Param		status	query	model.Status	false	"Filter by status"
//	@Success	200		{array}	model.Message
//	@Router		/messages [get]
func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	messageStatus := model.Status(c.Query("status"))
	messages, err := h.messageService.GetMessages(messageStatus)
	if err != nil {
		if errors.Is(err, model.ErrorInvalidMessageStatus) { // TODO: put here to middleware
			return fiber.ErrBadRequest
		}

		if errors.Is(err, model.ErrorMessageNotFound) {
			return fiber.ErrNotFound
		}

		return err
	}

	return c.JSON(messages)
}
