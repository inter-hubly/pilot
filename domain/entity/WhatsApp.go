package entity

type whatsAppMessage struct {
	Body string `json:"body"`
}

func NewWhatsAppMessage(body string) *whatsAppMessage {
	return &whatsAppMessage{
		Body: body,
	}

}

func (w *whatsAppMessage) GetBody() string {
	return w.Body
}

func NormalizeWhatsAppMessage(message *WhatsAppJSONReceived) *Chat {
	return &Chat{
		Id:        message.Id,
		MessageId: message.Metadata.MessageId,
		Message:   NewWhatsAppMessage(message.Metadata.Body),
		Delivered: ChatMessageTime{
			Status:     "sending",
			ReceivedAt: message.Metadata.Timestamp,
		},
	}
}
