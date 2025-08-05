package model_test

import (
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestStatus_IsValid(t *testing.T) {
	testCases := []struct {
		description    string
		expectedResult bool
		status         string
	}{
		{
			description:    "unsent status",
			expectedResult: true,
			status:         "UNSENT",
		},
		{
			description:    "sent status",
			expectedResult: true,
			status:         "SENT",
		},
		{
			description:    "rejected status",
			expectedResult: true,
			status:         "REJECTED",
		},
		{
			description:    "invalid status",
			expectedResult: false,
			status:         "INVALID",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualResult := model.Status(testCase.status).IsValid()
			assert.Equal(t, testCase.expectedResult, actualResult)
		})
	}
}
