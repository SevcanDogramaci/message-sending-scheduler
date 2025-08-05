package client_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	client "github.com/SevcanDogramaci/message-sending-scheduler/internal/client/webhook_site"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/stretchr/testify/assert"
)

const testWebhookSiteAPIKey = "test-api-key"

func TestWebhookClient_GivenUnsentMessage_ThenItShouldSendIt(t *testing.T) {
	message := model.Message{
		ID:               "test-msg-id",
		Content:          "test-msg-content",
		SenderPhoneNo:    "+901111111111",
		RecipientPhoneNo: "+902222222222",
		Status:           model.StatusUnsent,
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		actualBody, _ := io.ReadAll(req.Body)
		defer req.Body.Close()

		expectedBody := `{"to":"+902222222222","content":"test-msg-content"}`

		assert.Equal(t, http.MethodPost, req.Method)
		assert.JSONEq(t, expectedBody, string(actualBody))
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, testWebhookSiteAPIKey, req.Header.Get("x-ins-auth-key"))

		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte(`{"message": "Accepted","messageId": "02aa861c-27d4-4f0e-a77b-3720794376e8"}`))
	}))
	defer server.Close()

	client := client.NewWebhookSiteClient(&config.ClientConfig{
		URL:    server.URL,
		APIKey: testWebhookSiteAPIKey,
	})

	err := client.Send(message)
	assert.NoError(t, err)
}

func TestWebhookClient_GivenUnsuccessfulResponse_ThenItShouldReturnError(t *testing.T) {
	message := model.Message{}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := client.NewWebhookSiteClient(&config.ClientConfig{
		URL:    server.URL,
		APIKey: testWebhookSiteAPIKey,
	})

	err := client.Send(message)
	assert.Error(t, err)
}
