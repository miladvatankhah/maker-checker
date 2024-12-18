package use_cases

import (
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/repositories"
)

type RejectMessageUseCase struct {
	MessageRepo repositories.MessageRepository
}

func NewRejectMessageUseCase(messageRepo repositories.MessageRepository) *RejectMessageUseCase {
	return &RejectMessageUseCase{MessageRepo: messageRepo}
}

func (uc *RejectMessageUseCase) RejectMessage(id string) error {
	// Retrieve the message from the repository
	message, err := uc.MessageRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Update the message status
	message.Reject()

	// Save the updated message status in the repository
	return uc.MessageRepo.Save(message)
}
