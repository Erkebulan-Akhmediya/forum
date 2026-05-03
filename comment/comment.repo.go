package comment

import "forum/db"

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) savePost(c *postComment) error {
	query := `insert into comment (content, author_id, post_id) 
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, c.content, c.authorId, c.postId)
	return row.Scan(&c.id)
}
