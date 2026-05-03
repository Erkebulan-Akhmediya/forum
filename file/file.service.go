package file

type Service struct {
	repo *repo
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) UploadPost(name string, postId int) error {
	file := PostFile{
		File:   File{Name: name},
		PostId: postId,
	}
	return s.repo.savePost(&file)
}
