package use_cases

import (
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/mappers"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/repositories"
)

type RegisterUserUseCase struct {
	UserRepo repositories.UserRepository
}

func NewRegisterUserUseCase(userRepo repositories.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{UserRepo: userRepo}
}

func (uc *RegisterUserUseCase) RegisterUser(dto dtos.RegisterUserDtoRequest) error {
	// Map DTO to domain entity
	user := mappers.ToDomainUser(dto)

	// Save user in repository
	return uc.UserRepo.Save(user)
}
