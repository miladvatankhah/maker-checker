package dtos

import "github.com/google/uuid"

type RegisterUserDtoRequest struct {
	ID uuid.UUID `json:"id"`
}
