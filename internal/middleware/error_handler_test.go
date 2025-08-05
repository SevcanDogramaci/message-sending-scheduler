package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/middleware"
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	testCases := []struct {
		description        string
		err                error
		expectedStatusCode int
	}{
		{
			description:        "given invalid error, then should return bad request",
			err:                model.ErrorInvalidMessageStatus,
			expectedStatusCode: 400,
		},
		{
			description:        "given not found error, then should return not found",
			err:                model.ErrorMessageNotFound,
			expectedStatusCode: 404,
		},
		{
			description:        "given unexpected error, then should return internal server error",
			err:                model.ErrorMessageStatusNotUpdated,
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			url := "/test"
			app := fiber.New(fiber.Config{
				ErrorHandler: middleware.InitErrorHandler,
			})

			app.Get(url, func(c *fiber.Ctx) error {
				return testCase.err
			})

			req := httptest.NewRequest(http.MethodGet, url, nil)
			resp, err := app.Test(req, -1)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedStatusCode, resp.StatusCode)
		})
	}
}
