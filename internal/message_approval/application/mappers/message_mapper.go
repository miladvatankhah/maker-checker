package mappers

import (
	"github.com/miladvatankhah/maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/entities"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/value_objects"
)

func ToDTOMessage(message *entities.Message) dtos.CreateMessageRequestDTO {
	return dtos.CreateMessageRequestDTO{
		Content:    message.Content.Text,
		SenderID:   message.SenderID,
		ReceiverID: message.ReceiverID,
	}
}

func ToDomainMessage(dto dtos.CreateMessageRequestDTO) *entities.Message {
	return &entities.Message{
		Content:    value_objects.MessageContent{Text: dto.Content},
		SenderID:   dto.SenderID,
		ReceiverID: dto.ReceiverID,
	}
}
