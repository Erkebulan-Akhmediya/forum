package post

import "mime/multipart"

type createDto struct {
	title, content string
	files          []*multipart.FileHeader
	authorId       int
}
