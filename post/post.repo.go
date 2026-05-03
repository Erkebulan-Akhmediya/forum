package post

import "forum/db"

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) save(p *post) error {
	query := `insert into post (title, content, author_id) 
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, p.title, p.content, p.authorId)
	err := row.Scan(&p.id)
	return err
}
