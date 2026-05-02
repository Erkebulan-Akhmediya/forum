package auth

import (
	"encoding/json"
	"forum/utils"
	"log"
	"net/http"
	"slices"
)

type signUpHandler struct {
	service *service
}

func NewSignUpHandler() http.Handler {
	return &signUpHandler{
		service: newService(),
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

	err := h.service.createUser(dto.Username, dto.Password, dto.Email)
	if err != nil {
		code := 500
		userErrs := []error{errDuplicateEmail, errPwdTooLong}
		if slices.Contains(userErrs, err) {
			code = 400
		}
		utils.SendMessage(w, err.Error(), code)
		return
	}
	utils.SendMessage(w, "You have successfully sign up!", 201)
}
