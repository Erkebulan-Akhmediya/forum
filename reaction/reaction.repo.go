package reaction

import "forum/db"

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (repo *repo) savePostReaction(r *postReaction) error {
	query := `insert into reaction (user_id, post_id, reaction) 
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, r.userId, r.postId, r.reaction.reaction)
	return row.Scan(&r.id)
}

func (repo *repo) saveCommentReaction(r *commentReaction) error {
	query := `insert into reaction (user_id, comment_id, reaction)
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, r.userId, r.commentId, r.reaction.reaction)
	return row.Scan(&r.id)
}
