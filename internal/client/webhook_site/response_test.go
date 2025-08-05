package client_test

import (
	"testing"
	"time"

	client "github.com/SevcanDogramaci/message-sending-scheduler/internal/client/webhook_site"
	"github.com/stretchr/testify/assert"
)

func TestMessageResponse_ToTransferMetadata(t *testing.T) {
	messageResponse := client.MessageResponse{MessageID: "test-message-id"}
	metadata := messageResponse.ToTransferMetadata()

	assert.Equal(t, messageResponse.MessageID, metadata.ID)
	assert.NotEqual(t, time.Time{}, metadata.TransferTime)
}
