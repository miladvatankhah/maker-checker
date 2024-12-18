package events

import "github.com/google/uuid"

type MessageApprovedEvent struct {
	MessageID  uuid.UUID
	ReceiverID uuid.UUID
}
