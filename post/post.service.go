package post

import (
	"forum/file"
	"forum/utils"
)

type service struct {
	repo        *repo
	fileService *file.Service
}

func newService() *service {
	return &service{
		repo:        newRepo(),
		fileService: file.NewService(),
	}
}

func (s *service) create(dto *createDto) error {
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

func (s *service) getAll(page *utils.Page) ([]*post, error) {
	return s.repo.getAll(page.Index, page.Size)
}
