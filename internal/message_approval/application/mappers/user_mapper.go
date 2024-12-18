package mappers

import (
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/aggregates"
)

func ToDTOUser(user *aggregates.User) dtos.RegisterUserDtoRequest {
	return dtos.RegisterUserDtoRequest{
		ID: user.ID,
	}
}

func ToDomainUser(dto dtos.RegisterUserDtoRequest) *aggregates.User {
	return &aggregates.User{
		ID: dto.ID,
	}
}
