package auth

import (
	"encoding/json"
	"forum/user"
	"forum/utils"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

	pwdBytes := []byte(dto.Password)
	// this check is necessary because bcrypt.GenerateFromPassword
	// does not operate on byte arrays longer that 72
	if len(pwdBytes) > 72 {
		utils.SendMessage(w, "Password too long", 400)
		return
	}

	pwdBytes, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		utils.SendMessage(w, "Failed to hash password", 500)
		return
	}

	if err := h.service.Create(dto.Username, string(pwdBytes), dto.Email); err != nil {
		log.Println("Error creating new user:", err)
		utils.SendMessage(w, "Sign up failed", 500)
		return
	}
	utils.SendMessage(w, "You have successfully sign up!", 201)
}
