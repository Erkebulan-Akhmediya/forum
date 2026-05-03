package comment

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

func (s *service) createPostComment(dto *createPostCommentDto) error {
	c := postComment{
		comment: comment{
			content:  dto.content,
			authorId: dto.auhtorId,
		},
		postId: dto.postId,
	}
	if err := s.repo.savePostComment(&c); err != nil {
		return err
	}
	return s.fileService.UploadCommentFile(c.id, dto.file)
}
