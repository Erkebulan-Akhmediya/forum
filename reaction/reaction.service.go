package reaction

const (
	like    = "LIKE"
	dislike = "DISLIKE"
)

type Service struct {
	repo *repo
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) LikePost(userId, postId int) error {
	return s.createPostReaction(userId, postId, like)
}

func (s *Service) DislikePost(userId, postId int) error {
	return s.createPostReaction(userId, postId, dislike)
}

func (s *Service) createPostReaction(userId, postId int, reactionType string) error {
	r := postReaction{
		reaction: reaction{
			userId:   userId,
			reaction: reactionType,
		},
		postId: postId,
	}
	return s.repo.savePostReaction(&r)
}

func (s *Service) LikeComment(userId, commentId int) error {
	return s.createCommentReaction(userId, commentId, like)
}

func (s *Service) DislikeComment(userId, commentId int) error {
	return s.createCommentReaction(userId, commentId, dislike)
}

func (s *Service) createCommentReaction(userId, commentId int, reactionType string) error {
	r := commentReaction{
		reaction: reaction{
			userId:   userId,
			reaction: reactionType,
		},
		commentId: commentId,
	}
	return s.repo.saveCommentReaction(&r)
}
