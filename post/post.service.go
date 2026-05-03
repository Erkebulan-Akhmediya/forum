package post

import "forum/file"

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

func (s *service) create(dto *createPostDto) error {
	p := post{
		title:    dto.title,
		content:  dto.content,
		authorId: dto.authorId,
	}
	if err := s.repo.save(&p); err != nil {
		return err
	}
	for _, fh := range dto.files {
		err := s.fileService.UploadPost(fh.Filename, p.id)
		if err != nil {
			return err
		}
	}
	return nil
}
