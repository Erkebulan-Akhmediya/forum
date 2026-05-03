package post

type post struct {
	id, authorId   int
	title, content string
	author         author
}

type author struct {
	id              int
	username, email string
}
