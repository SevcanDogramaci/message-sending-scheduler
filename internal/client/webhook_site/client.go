package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SevcanDogramaci/message-sending-scheduler/config"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
)

type WebhookSiteClient struct {
	url    string
	apiKey string
}

func NewWebhookSiteClient(config *config.ClientConfig) *WebhookSiteClient {
	return &WebhookSiteClient{url: config.URL, apiKey: config.APIKey}
}

type MessageDTO struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

func (c *WebhookSiteClient) Send(message model.Message) error {
	messageDTO := MessageDTO{
		To:      message.RecipientPhoneNo,
		Content: message.Content,
	}

	messageJSON, err := json.Marshal(messageDTO)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-ins-auth-key", c.apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error occurred while sending request - Body: %s", responseBody)
	}

	return nil
}
