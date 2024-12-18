package use_cases

import (
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/events"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/domain/repositories"
)

type ApproveMessageUseCase struct {
	MessageRepo    repositories.MessageRepository
	EventPublisher events.DomainEventPublisher
}

func NewApproveMessageUseCase(messageRepo repositories.MessageRepository, eventPublisher events.DomainEventPublisher) *ApproveMessageUseCase {
	return &ApproveMessageUseCase{MessageRepo: messageRepo, EventPublisher: eventPublisher}
}

func (uc *ApproveMessageUseCase) ApproveMessage(id string) error {
	// Retrieve the message from the repository
	message, err := uc.MessageRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Update the message status
	message.Approve()

	// Save the updated message status in the repository
	if err := uc.MessageRepo.Save(message); err != nil {
		return err
	}

	// Create and publish the approved event
	event := events.MessageApprovedEvent{
		MessageID:  message.ID,
		ReceiverID: message.ReceiverID,
	}
	uc.EventPublisher.Publish(event)

	return nil
}
