package auth

import (
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	signUpHandler := utils.MethodHandler{
		http.MethodPost: NewSignUpHandler(),
	}
	http.Handle("/auth/sign-up", signUpHandler)
}
