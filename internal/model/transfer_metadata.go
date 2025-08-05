package model

import "time"

type TransferMetadata struct {
	ID           string    `json:"id"`
	TransferTime time.Time `json:"transferTime"`
}
