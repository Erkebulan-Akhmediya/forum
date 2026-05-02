package post

import "mime/multipart"

type createPostDto struct {
	title, content string
	files          []*multipart.FileHeader
	authorId       int
}
