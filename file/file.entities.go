package file

type File struct {
	Name string
	Id   int
}

type PostFile struct {
	File
	PostId int
}

type CommentFile struct {
	File
	CommentId int
}
