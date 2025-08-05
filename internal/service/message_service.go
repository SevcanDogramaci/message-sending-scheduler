package service

import (
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2/log"
)

type CacheRepository interface {
	SetMessage(metadata model.TransferMetadata) error
}

type MessageRepository interface {
	GetMessagesByStatus(status model.Status, limit int) ([]model.Message, error)
	UpdateMessageStatus(msg model.Message, status model.Status) (model.Message, error)
}

type MessageClient interface {
	Send(message model.Message) (*model.TransferMetadata, error)
}

type MessageService struct {
	client     MessageClient
	repository MessageRepository
	cache      CacheRepository
}

func NewMessageService(client MessageClient, repository MessageRepository, cache CacheRepository) *MessageService {
	return &MessageService{
		client:     client,
		repository: repository,
		cache:      cache,
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

		transferMetadata, err := ms.client.Send(msg)
		if err != nil {
			return err
		}

		msg, err = ms.repository.UpdateMessageStatus(msg, model.StatusSent)
		if err != nil {
			log.Error("[ACTION REQUIRED] Failed to update message status:", err)
			return model.ErrorMessageStatusNotUpdated
		}

		err = ms.cache.SetMessage(*transferMetadata)
		if err != nil {
			log.Errorf("Failed to cache message with transfer id: %s", transferMetadata.ID)
			return model.ErrorMessageTransferNotCached
		}
	}

	return nil
}
