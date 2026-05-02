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

func newSignUpHandler() http.Handler {
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

type signInHandler struct {
	service *service
}

func newSignInHandler() http.Handler {
	return &signInHandler{
		service: newService(),
	}
}

func (h *signInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dto signInDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println("Error parsing sign in request:", err)
		utils.SendMessage(w, "Error parsing request", 400)
		return
	}

	userId, err := h.service.validateCredentials(dto.Email, dto.Password)
	if err != nil {
		log.Println("Error validating credentials:", err)
		utils.SendMessage(w, "Invalid credentials", 400)
		return
	}

	sid, err := h.service.createSession(userId)
	if err != nil {
		log.Println("Error creating session:", err)
		utils.SendMessage(w, "Failed to create session", 500)
		return
	}

	cookie := http.Cookie{
		Name:  "cookie",
		Value: sid,
	}
	http.SetCookie(w, &cookie)
	utils.SendMessage(w, "You have successfully signed in!", 200)
}
