package comment

import (
	"mime/multipart"
)

type createPostCommentDto struct {
	content  string
	auhtorId int
	postId   int
	file     *multipart.FileHeader
}

type getPostCommentDto struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Author  authorDto `json:"author"`
	FileId  int       `json:"fileId,omitempty"`
}

type authorDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
