package post

import (
	"forum/file"
	"forum/utils"
)

type Service struct {
	repo        *repo
	fileService *file.Service
}

func NewService() *Service {
	return &Service{
		repo:        newRepo(),
		fileService: file.NewService(),
	}
}

func (s *Service) create(dto *createDto) error {
	p := post{
		title:    dto.title,
		content:  dto.content,
		authorId: dto.authorId,
	}
	if err := s.repo.save(&p); err != nil {
		return err
	}
	for _, fh := range dto.files {
		if err := s.fileService.UploadPost(p.id, fh); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) ExistsById(id int) (bool, error) {
	return s.repo.existsById(id)
}

func (s *Service) getAll(page *utils.Page) ([]*post, error) {
	return s.repo.getAll(page.Index, page.Size)
}
