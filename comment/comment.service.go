package comment

import (
	"errors"
	"forum/file"
	"forum/post"
	"forum/utils"
	"log"
)

var errCheckPost = errors.New("Failed to check if post exists")
var errPostNotFound = errors.New("Post not found")
var errCheckComment = errors.New("Failed to check if comment exists")
var errCommentNotFound = errors.New("Comment not found")

type service struct {
	repo        *repo
	fileService *file.Service
	postService *post.Service
}

func newService() *service {
	return &service{
		repo:        newRepo(),
		fileService: file.NewService(),
		postService: post.NewService(),
	}
}

func (s *service) createPostComment(dto *createPostCommentDto) error {
	exists, err := s.postService.ExistsById(dto.postId)
	if err != nil {
		log.Println("Error checking if post exists:", err)
		return errCheckPost
	}
	if !exists {
		return errPostNotFound
	}

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

func (s *service) createReplyComment(dto *createReplyCommentDto) error {
	exists, err := s.repo.existsById(dto.commentId)
	if err != nil {
		return errCheckComment
	}
	if !exists {
		return errCommentNotFound
	}
	c := replyComment{
		comment: comment{
			content:  dto.content,
			authorId: dto.authorId,
		},
		commentId: dto.commentId,
	}
	if err := s.repo.saveReplyComment(&c); err != nil {
		return err
	}
	return s.fileService.UploadCommentFile(c.id, dto.file)
}

func (s *service) getAllByPostId(postId int, page *utils.Page) ([]*postComment, error) {
	return s.repo.getAllByPostId(postId, page)
}
