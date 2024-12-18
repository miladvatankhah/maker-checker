package v1

import (
	"encoding/json"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/application/dtos"
	"github.com/miladvatankhah/maker-checker/internal/message_approval/application/use_cases"
	"net/http"
)

type UserHandler struct {
	RegisterUserUseCase *use_cases.RegisterUserUseCase
}

func NewUserHandler(uc *use_cases.RegisterUserUseCase) *UserHandler {
	return &UserHandler{RegisterUserUseCase: uc}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dtos.RegisterUserDtoRequest
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.RegisterUserUseCase.RegisterUser(userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
