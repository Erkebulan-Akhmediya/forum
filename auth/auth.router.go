package auth

import (
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	signUpHandler := utils.MethodHandler{
		http.MethodPost: newSignUpHandler(),
	}
	http.Handle("/auth/sign-up", signUpHandler)

	signInHandler := utils.MethodHandler{
		http.MethodPost: newSignInHandler(),
	}
	http.Handle("/auth/sign-in", signInHandler)
}
