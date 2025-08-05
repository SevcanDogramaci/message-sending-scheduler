package service

import (
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2/log"
)

type MessageRepository interface {
	GetMessagesByStatus(status model.Status, limit int) ([]model.Message, error) // add limit
	UpdateMessageStatus(msg model.Message, status model.Status) (model.Message, error)
}

type MessageClient interface {
	Send(message model.Message) error
}

type MessageService struct {
	repository MessageRepository
	client     MessageClient
}

func NewMessageService(repository MessageRepository, client MessageClient) *MessageService {
	return &MessageService{
		repository: repository,
		client:     client,
	}
}

const SendMessageLimit = 2
const DefaultMessageLimit = 1000

func (ms *MessageService) GetMessages(status model.Status) ([]model.Message, error) {
	if !status.IsValid() {
		return nil, model.ErrorInvalidMessageStatus
	}

	messages, err := ms.repository.GetMessagesByStatus(status, DefaultMessageLimit)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, model.ErrorMessageNotFound
	}

	return messages, nil
}

func (ms *MessageService) SendMessages() error {
	messages, err := ms.repository.GetMessagesByStatus(model.StatusUnsent, SendMessageLimit)
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		return model.ErrorMessageNotFound
	}

	for _, msg := range messages {
		if !msg.IsValid() {
			_, err := ms.repository.UpdateMessageStatus(msg, model.StatusRejected)
			if err != nil {
				log.Error("[ACTION REQUIRED] Failed to update message status:", err)
				return model.ErrorMessageStatusNotUpdated
			}

			continue
		}

		err := ms.client.Send(msg)
		if err != nil {
			return err
		}

		msg, err = ms.repository.UpdateMessageStatus(msg, model.StatusSent)
		if err != nil {
			log.Error("[ACTION REQUIRED] Failed to update message status:", err)
			return model.ErrorMessageStatusNotUpdated
		}
	}

	return nil
}
