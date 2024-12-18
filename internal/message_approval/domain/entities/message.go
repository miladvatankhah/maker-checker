package entities

import (
	"github.com/google/uuid"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/value_objects"
)

type MessageStatus string

const (
	Pending  MessageStatus = "Pending"
	Approved MessageStatus = "Approved"
	Rejected MessageStatus = "Rejected"
)

type Message struct {
	ID         uuid.UUID
	Content    value_objects.MessageContent
	Status     MessageStatus
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
}

func NewMessage(content string, senderID, receiverID uuid.UUID) *Message {
	return &Message{
		ID:         generateID(), // Implement ID generation logic
		Content:    value_objects.MessageContent{Text: content},
		Status:     Pending,
		SenderID:   senderID,
		ReceiverID: receiverID,
	}
}

func (m *Message) Approve() {
	m.Status = Approved
}

func (m *Message) Reject() {
	m.Status = Rejected
}

func generateID() uuid.UUID {
	return uuid.New()
}
