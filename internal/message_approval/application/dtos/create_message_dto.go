package dtos

import "github.com/google/uuid"

type CreateMessageRequestDTO struct {
	Content    string    `json:"content"`
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
}
