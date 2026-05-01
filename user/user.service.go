package user

type Service struct {
	repo *repo
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) Create(username, password, email string) error {
	user := User{
		Username: username,
		Password: password,
		Email:    email,
	}
	return s.repo.save(&user)
}
