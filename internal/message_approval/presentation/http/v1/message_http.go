package v1

import (
	"encoding/json"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/application/use_cases"
	"net/http"
)

type MessageHandler struct {
	CreateMessageUseCase  *use_cases.CreateMessageUseCase
	ApproveMessageUseCase *use_cases.ApproveMessageUseCase
	RejectMessageUseCase  *use_cases.RejectMessageUseCase
}

func NewMessageHandler(createUC *use_cases.CreateMessageUseCase, approveUC *use_cases.ApproveMessageUseCase, rejectUC *use_cases.RejectMessageUseCase) *MessageHandler {
	return &MessageHandler{
		CreateMessageUseCase:  createUC,
		ApproveMessageUseCase: approveUC,
		RejectMessageUseCase:  rejectUC,
	}
}

func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var messageDTO dtos.CreateMessageRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&messageDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.CreateMessageUseCase.CreateMessage(messageDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *MessageHandler) ApproveMessage(w http.ResponseWriter, r *http.Request) {
	msgID := r.PathValue("id")
	if err := h.ApproveMessageUseCase.ApproveMessage(msgID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *MessageHandler) RejectMessage(w http.ResponseWriter, r *http.Request) {
	msgID := r.PathValue("id")
	if err := h.RejectMessageUseCase.RejectMessage(msgID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
