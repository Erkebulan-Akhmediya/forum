package auth

import (
	"encoding/json"
	"forum/user"
	"forum/utils"
	"log"
	"net/http"
)

type signUpHandler struct {
	service *user.Service
}

func NewSignUpHandler() http.Handler {
	return &signUpHandler{
		service: user.NewService(),
	}
}

func (h *signUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dto signUpDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println("Error parsing sign up request:", err)
		utils.SendMessage(w, "Error parsing request", 400)
		return
	}
	if err := h.service.Create(dto.Username, dto.Password, dto.Email); err != nil {
		log.Println("Error creating new user:", err)
		utils.SendMessage(w, "Sign up failed", 500)
		return
	}
	utils.SendMessage(w, "You have successfully sign up!", 201)
}
