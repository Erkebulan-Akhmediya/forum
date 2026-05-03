package comment

type postComment struct {
	comment
	postId int
}

type replyComment struct {
	comment
	commentId int
}

type comment struct {
	id       int
	content  string
	authorId int
	author   author
}

type author struct {
	id              int
	username, email string
}
