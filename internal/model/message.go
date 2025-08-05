package model

const MessageCharLimit = 10

type Message struct {
	ID               string `json:"id"`
	Content          string `json:"content"`
	SenderPhoneNo    string `json:"sender_phone_no"`
	RecipientPhoneNo string `json:"recipient_phone_no"`
	Status           Status `json:"status"`
}

func (m Message) IsValid() bool {
	return len(m.Content) <= MessageCharLimit
}
