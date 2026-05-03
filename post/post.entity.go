package post

type post struct {
	id, authorId   int
	title, content string
	author         author
	fileIds        []int
}

type author struct {
	id              int
	username, email string
}
