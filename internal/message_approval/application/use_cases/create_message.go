package use_cases

import (
	"errors"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/repositories"
	"log"
)

type CreateMessageUseCase struct {
	UserRepo    repositories.UserRepository
	MessageRepo repositories.MessageRepository
}

func NewCreateMessageUseCase(userRepo repositories.UserRepository, messageRepo repositories.MessageRepository) *CreateMessageUseCase {
	return &CreateMessageUseCase{UserRepo: userRepo, MessageRepo: messageRepo}
}

func (uc *CreateMessageUseCase) CreateMessage(dto dtos.CreateMessageRequestDTO) error {
	// Retrieve sender and receiver from the user repository
	log.Println(dto.SenderID)
	sender, err := uc.UserRepo.FindByID(dto.SenderID)
	if err != nil {
		return errors.New("sender not found")
	}

	receiver, err := uc.UserRepo.FindByID(dto.ReceiverID)
	if err != nil {
		return errors.New("receiver not found")
	}

	// Create and save message
	message := sender.SendMessage(dto.Content, receiver)

	// Save message in repository
	return uc.MessageRepo.Save(message)
}
