package comment

import "mime/multipart"

type createPostCommentDto struct {
	content  string
	auhtorId int
	postId   int
	file     *multipart.File
}
