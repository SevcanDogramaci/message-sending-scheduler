package model_test

import (
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMessage_IsValid(t *testing.T) {
	testCases := []struct {
		description    string
		msg            model.Message
		expectedResult bool
	}{
		{
			description:    "short content",
			msg:            model.Message{Content: "valid"},
			expectedResult: true,
		},
		{
			description: "long content",
			msg: model.Message{
				Content: "veryveryverylongcontentveryveryverylongcontentveryveryverylongcontent"},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.msg.IsValid())
		})
	}
}
