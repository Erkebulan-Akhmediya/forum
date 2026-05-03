package post

import "mime/multipart"

type createDto struct {
	title, content string
	files          []*multipart.FileHeader
	authorId       int
}

type getDto struct {
	Id      int       `json:"id"`
	Author  authorDto `json:"author"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	FileIds []int     `json:"fileIds"`
}

type authorDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
