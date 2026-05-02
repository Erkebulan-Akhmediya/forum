package auth

type signUpDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type signInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
