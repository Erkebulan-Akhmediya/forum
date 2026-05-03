package file

import "forum/db"

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) savePost(file *PostFile) error {
	query := "insert into file (name, post_id) values (?, ?)"
	_, err := db.DB.Exec(query, file.Name, file.PostId)
	return err
}

func (r *repo) saveComment(file *CommentFile) error {
	query := "insert into file (name, comment_id) values (?, ?)"
	_, err := db.DB.Exec(query, file.Name, file.CommentId)
	return err
}
