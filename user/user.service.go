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

func (s *Service) ExistsByEmail(email string) (bool, error) {
	return s.repo.existsByEmail(email)
}

func (s *Service) GetByEmail(email string) (*User, error) {
	return s.repo.getByEmail(email)
}
