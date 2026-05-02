package auth

import (
	"context"
	"encoding/json"
	"forum/utils"
	"log"
	"net/http"
	"slices"
	"time"
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

	s, err := h.service.getSessionByUserId(userId)
	if err != nil {
		log.Println("Error getting session:", err)
	} else {
		sendSidCookie(w, s)
		utils.SendMessage(w, "You have successfully signed in!", 200)
		return
	}

	s, err = h.service.createSession(userId)
	if err != nil {
		log.Println("Error creating session:", err)
		utils.SendMessage(w, "Failed to create session", 500)
		return
	}

	sendSidCookie(w, s)
	utils.SendMessage(w, "You have successfully signed in!", 200)
}

func sendSidCookie(w http.ResponseWriter, s *session) {
	cookie := http.Cookie{
		Name:       "sid",
		Value:      s.id,
		Path:       "/",
		Expires:    s.expiresAt,
		RawExpires: s.expiresAt.String(),
	}
	http.SetCookie(w, &cookie)
}

type authMiddleware struct {
	next    http.Handler
	service *service
}

func NewMiddleware(next http.Handler) http.Handler {
	return &authMiddleware{
		next:    next,
		service: newService(),
	}
}

func (m *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sid")
	if err != nil {
		log.Println("Error in auth middleware:", err)
		utils.SendMessage(w, "Cookie not found", 401)
		return
	}
	s, err := m.service.getSessionById(cookie.Value)
	if err != nil {
		log.Println("Error in auth middleware:", err)
		utils.SendMessage(w, "Session not found", 401)
		return
	}
	if time.Now().After(s.expiresAt) {
		utils.SendMessage(w, "Session expired", 401)
		return
	}
	ctx := context.WithValue(r.Context(), "userId", s.userId)
	m.next.ServeHTTP(w, r.WithContext(ctx))
}
