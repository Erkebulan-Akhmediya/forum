package reaction

const (
	like    = "LIKE"
	dislike = "DISLIKE"
)

type service struct {
	repo *repo
}

func newService() *service {
	return &service{
		repo: newRepo(),
	}
}

func (s *service) LikePost(userId, postId int) error {
	return s.createPostReaction(userId, postId, like)
}

func (s *service) DislikePost(userId, postId int) error {
	return s.createPostReaction(userId, postId, dislike)
}

func (s *service) createPostReaction(userId, postId int, reactionType string) error {
	r := postReaction{
		reaction: reaction{
			userId:   userId,
			reaction: reactionType,
		},
		postId: postId,
	}
	return s.repo.savePostReaction(&r)
}

func (s *service) LikeComment(userId, postId int) error {
	return s.createCommentReaction(userId, postId, like)
}

func (s *service) DislikeComment(userId, postId int) error {
	return s.createCommentReaction(userId, postId, dislike)
}

func (s *service) createCommentReaction(userId, commentId int, reactionType string) error {
	r := commentReaction{
		reaction: reaction{
			userId:   userId,
			reaction: reactionType,
		},
		commentId: commentId,
	}
	return s.repo.saveCommentReaction(&r)
}
