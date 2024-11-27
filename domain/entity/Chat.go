package entity

type Chat struct {
	MessageId string                  `json:"messageId"`
	OwnerId   string                  `json:"ownerId"`
	From      string                  `json:"from"`
	To        string                  `json:"to"`
	Audit     []ChatMessageStatusTime `json:"status,omitempty"`
	Message   ReceivedMessage         `json:"message,omitempty"`
}

type ReceivedMessage interface {
	GetBody() string
}

type ChatMessageStatusTime struct {
	Status     MessageStatus `json:"status"`
	ReceivedAt string        `json:"receivedAt"`
}
