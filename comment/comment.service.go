package comment

type service struct {
	repo *repo
}

func newService() *service {
	return &service{
		repo: newRepo(),
	}
}

func (s *service) createPost(dto *createPostCommentDto) error {
	c := postComment{
		comment: comment{
			content:  dto.content,
			authorId: dto.auhtorId,
		},
		postId: dto.postId,
	}
	return s.repo.savePost(&c)
}
