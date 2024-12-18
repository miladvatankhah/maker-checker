package repositories

import "github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/entities"

type MessageRepository interface {
	Save(message *entities.Message) error
	FindByID(id string) (*entities.Message, error)
}