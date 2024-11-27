package entity

type ChatType string

const WhatsApp ChatType = "WhatsApp"

type Chat struct {
	Id        string          `json:"id"`
	MessageId string          `json:"messageId"`
	Type      ChatType        `json:"type"`
	Status    ChatMessageTime `json:"received,omitempty"`
	Read      ChatMessageTime `json:"read,omitempty"`
	Delivered ChatMessageTime `json:"delivered,omitempty"`
	Message   ReceivedMessage `json:"message,omitempty"`
}

type ReceivedMessage interface {
	GetBody() string
}

type ChatMessageTime struct {
	Status     string `json:"status"`
	ReceivedAt string `json:"receivedAt"`
}
