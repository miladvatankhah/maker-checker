package repositories

import (
	"github.com/google/uuid"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/aggregates"
)

type UserRepository interface {
	Save(user *aggregates.User) error
	FindByID(id uuid.UUID) (*aggregates.User, error)
}
