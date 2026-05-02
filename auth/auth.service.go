package auth

import (
	"errors"
	"forum/user"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var errCheckEmail = errors.New("Failed to check if the email is unique")
var errDuplicateEmail = errors.New("The email is already in use")
var errPwdTooLong = errors.New("Password too long")
var errPwdHash = errors.New("Failed to hash password")
var errCreateUser = errors.New("Failed to create new user")

type service struct {
	userUservice *user.Service
	sessionRepo  *sessionRepo
}

func newService() *service {
	return &service{
		userUservice: user.NewService(),
		sessionRepo:  newSessionRepo(),
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

// validates credentials and returns user id if valid, errors out otherwise
func (s *service) validateCredentials(email, password string) (int, error) {
	user, err := s.userUservice.GetByEmail(email)
	if err != nil {
		return -1, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return -1, err
	}
	return user.Id, nil
}

// creates a sesssion and returns its id
func (s *service) createSession(userId int) (string, error) {
	id := uuid.NewString()
	epxiresAt := time.Now().Add(time.Hour * 24 * 6)
	ss := session{
		id:        id,
		userId:    userId,
		expiresAt: epxiresAt,
	}
	return id, s.sessionRepo.save(&ss)
}

func (s *service) getSessionById(id string) (*session, error) {
	return s.sessionRepo.getById(id)
}
