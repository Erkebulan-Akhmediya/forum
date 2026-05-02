package post

type service struct {
	repo *repo
}

func newService() *service {
	return &service{
		repo: newRepo(),
	}
}

func (s *service) create(title, content string, authorId int) error {
	p := post{
		title:    title,
		content:  content,
		authorId: authorId,
	}
	return s.repo.save(&p)
}
