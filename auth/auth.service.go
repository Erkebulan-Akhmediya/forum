package auth

import (
	"errors"
	"forum/user"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var errCheckEmail = errors.New("Failed to check if the email is unique")
var errDuplicateEmail = errors.New("The email is already in use")
var errPwdTooLong = errors.New("Password too long")
var errPwdHash = errors.New("Failed to hash password")
var errCreateUser = errors.New("Failed to create new user")

type service struct {
	userUservice *user.Service
}

func newService() *service {
	return &service{
		userUservice: user.NewService(),
	}
}

func (s *service) createUser(username, password, email string) error {
	exists, err := s.userUservice.ExistsByEmail(email)
	if err != nil {
		log.Println("Error checking if email unique:", err)
		return errCheckEmail
	}
	if exists {
		return errDuplicateEmail
	}

	pwdBytes := []byte(password)
	// this check is necessary because bcrypt.GenerateFromPassword
	// does not operate on byte arrays longer that 72
	if len(pwdBytes) > 72 {
		return errPwdTooLong
	}

	pwdBytes, err = bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errPwdHash
	}

	err = s.userUservice.Create(username, string(pwdBytes), email)
	if err != nil {
		log.Println("Error creating new user:", err)
		return errCreateUser
	}
	return nil
}
