package dto

type MessageType string

const (
	MessageTypeStatuses MessageType = "statuses"
	MessageTypeMessage  MessageType = "message"
)

type MessageStatus string

const (
	SentStatus      MessageStatus = "sent"
	DeliveredStatus MessageStatus = "delivered"
	ReadStatus      MessageStatus = "read"
	ReceivedStatus  MessageStatus = "received"
)

type WhatsAppJSONReceived struct {
	Id          string              `json:"id,omitempty"`
	MessageType MessageType         `json:"messageType"`
	Owner       WhatsAppPhoneIdDto  `json:"owner,omitempty"`
	SenderPhone string              `json:"senderPhone,omitempty"`
	Status      MessageStatus       `json:"status,omitempty"`
	Metadata    WhatsAppStatusesDto `json:"metadata,omitempty"`
}

type WhatsAppPhoneIdDto struct {
	PhoneNumberID      string `json:"phoneNumberId,omitempty"`
	DisplayPhoneNumber string `json:"displayPhoneNumber,omitempty"`
}

type WhatsAppStatusesDto struct {
	Timestamp      string `json:"timestamp,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	OriginType     string `json:"originType,omitempty"`
	MessageId      string `json:"messageId,omitempty"`
	Body           string `json:"body,omitempty"`
	BodyLength     int    `json:"bodyLength,omitempty"`
}
