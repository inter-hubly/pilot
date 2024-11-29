package entity

type Chat struct {
	MessageId string                  `json:"messageId"`
	OwnerId   string                  `json:"ownerId"`
	From      string                  `json:"from"`
	To        string                  `json:"to"`
	Audit     []ChatMessageStatusTime `json:"status,omitempty"`
	Message   string                  `json:"message,omitempty"`
}

type ChatMessageStatusTime struct {
	Status     MessageStatus `json:"status"`
	ReceivedAt string        `json:"receivedAt"`
}
