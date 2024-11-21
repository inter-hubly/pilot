package dto

type MessageType string

const (
	MessageTypeStatuses MessageType = "statuses"
	MessageTypeContacts MessageType = "contacts"
	MessageTypeMessage  MessageType = "message"
)

type WhatsAppJSONReceived struct {
	ID          string             `json:"id,omitempty"`
	MessageType MessageType        `json:"messageType"`
	Owner       WhatsAppPhoneIdDto `json:"sender,omitempty"`
	SenderPhone string             `json:"receive,omitempty"`
	Metadata    interface{}        `json:"metadata,omitempty"`
}

type WhatsAppPhoneIdDto struct {
	PhoneNumberID      string `json:"phoneNumberId,omitempty"`
	DisplayPhoneNumber string `json:"displayPhoneNumber,omitempty"`
}

// type WhatsAppMetadataDto struct {
// 	MessageID   string `json:"messageId,omitempty"`
// 	RecipientID string `json:"recipientId,omitempty"`
// 	Status      string `json:"status,omitempty"`
// 	Body        string `json:"body,omitempty"`
// 	Timestamp   int64  `json:"timestamp,omitempty"`
// 	BodyLength  int64  `json:"bodyLength,omitempty"`
// }

type WhatsAppStatusesDto struct {
	Timestamp      string `json:"timestamp,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	OriginType     string `json:"originType,omitempty"`
	MessageId      string `json:"messageId,omitempty"`
	Status         string `json:"status,omitempty"`
	Body           string `json:"body,omitempty"`
	BodyLength     int    `json:"bodyLength,omitempty"`
}
