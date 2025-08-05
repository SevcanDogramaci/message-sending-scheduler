package model

type Status string

const (
	StatusUnsent   Status = "UNSENT"
	StatusSent     Status = "SENT"
	StatusRejected Status = "REJECTED"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusUnsent, StatusSent, StatusRejected:
		return true
	default:
		return false
	}
}
