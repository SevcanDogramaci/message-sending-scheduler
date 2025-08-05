package model

import "errors"

var (
	ErrorInvalidMessageStatus    = errors.New("invalid message status")
	ErrorMessageNotFound         = errors.New("message not found")
	ErrorMessageStatusNotUpdated = errors.New("message status not updated")
)
