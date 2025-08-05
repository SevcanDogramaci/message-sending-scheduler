package service_test

import (
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/mocks"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMessages_GivenValidStatus_ThenItShouldReturnMessages(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	msgService := service.NewMessageService(mockRepository, nil)

	expectedMessages := []model.Message{{ID: "1", Status: model.StatusUnsent}}
	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusUnsent, service.DefaultMessageLimit).
		Return(expectedMessages, nil)

	actualMessages, err := msgService.GetMessages(model.StatusUnsent)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestGetMessages_GivenInvalidStatus_ThenItShouldReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	service := service.NewMessageService(nil, nil)
	actualMessages, err := service.GetMessages("INVALID")

	assert.Error(t, err)
	assert.Nil(t, actualMessages)
	assert.Equal(t, model.ErrorInvalidMessageStatus, err)
}

func TestGetMessages_GivenRepositoryError_ThenItShouldReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	msgService := service.NewMessageService(mockRepository, nil)

	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusUnsent, service.DefaultMessageLimit).
		Return(nil, assert.AnError)

	actualMessages, err := msgService.GetMessages(model.StatusUnsent)

	assert.Error(t, err)
	assert.Nil(t, actualMessages)
}

func TestGetMessages_GivenNoMessageFound_ThenItShouldReturnNotFoundError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	msgService := service.NewMessageService(mockRepository, nil)

	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusSent, service.DefaultMessageLimit).
		Return(nil, nil)

	actualMessages, err := msgService.GetMessages(model.StatusSent)

	assert.Error(t, err)
	assert.Nil(t, actualMessages)
}

func TestSendMessages_GivenUnsentMessages_ThenItShouldSendThem(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	mockClient := mocks.NewMockMessageClient(mockController)
	msgService := service.NewMessageService(mockRepository, mockClient)

	unsentMessages := []model.Message{
		{ID: "1", Status: model.StatusUnsent},
		{ID: "2", Status: model.StatusUnsent}}

	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusUnsent, service.SendMessageLimit).
		Return(unsentMessages, nil)

	for _, msg := range unsentMessages {
		mockClient.EXPECT().Send(msg).Return(nil)
		mockRepository.EXPECT().UpdateMessageStatus(msg, model.StatusSent).Return(msg, nil)
	}

	err := msgService.SendMessages()

	assert.NoError(t, err)
}

func TestSendMessages_GivenNoMessages_ThenItShouldReturnNotFoundError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	msgService := service.NewMessageService(mockRepository, nil)

	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusUnsent, service.SendMessageLimit).
		Return(nil, nil)

	err := msgService.SendMessages()

	assert.Equal(t, model.ErrorMessageNotFound, err)
}

func TestSendMessages_GivenMessageWithLongContent_ThenItShouldRejectTheMessage(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mocks.NewMockMessageRepository(mockController)
	mockClient := mocks.NewMockMessageClient(mockController)
	msgService := service.NewMessageService(mockRepository, mockClient)

	msgWithValidContent := model.Message{ID: "1", Status: model.StatusUnsent, Content: "proper"}
	msgWithInvalidContent := model.Message{ID: "2", Status: model.StatusUnsent, Content: "veryveryveryverylongcontent"}

	unsentMessages := []model.Message{msgWithValidContent, msgWithInvalidContent}

	mockRepository.
		EXPECT().
		GetMessagesByStatus(model.StatusUnsent, service.SendMessageLimit).
		Return(unsentMessages, nil)

	mockClient.EXPECT().Send(unsentMessages[0]).Return(nil)
	mockRepository.EXPECT().UpdateMessageStatus(msgWithValidContent, model.StatusSent).Return(msgWithValidContent, nil)
	mockRepository.EXPECT().UpdateMessageStatus(msgWithInvalidContent, model.StatusRejected).Return(msgWithInvalidContent, nil)

	err := msgService.SendMessages()

	assert.NoError(t, err)
}
