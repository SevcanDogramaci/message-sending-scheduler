package client

import (
	"time"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
)

type MessageResponse struct {
	MessageID string `json:"messageId"`
}

func (r *MessageResponse) ToTransferMetadata() *model.TransferMetadata {
	return &model.TransferMetadata{
		ID:           r.MessageID,
		TransferTime: time.Now(),
	}
}
