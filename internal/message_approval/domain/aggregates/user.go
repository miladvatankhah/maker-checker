package aggregates

import (
	"github.com/google/uuid"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/entities"
)

type User struct {
	ID               uuid.UUID
	SentMessages     []*entities.Message
	ReceivedMessages []*entities.Message
}

func NewUser(id uuid.UUID) *User {
	return &User{ID: id}
}

func (u *User) SendMessage(content string, receiver *User) *entities.Message {
	message := entities.NewMessage(content, u.ID, receiver.ID)
	u.SentMessages = append(u.SentMessages, message)
	receiver.ReceivedMessages = append(receiver.ReceivedMessages, message)
	message.Status = entities.Pending
	return message
}
