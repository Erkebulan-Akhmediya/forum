package post

import (
	"forum/db"
	"log"
)

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

func (r *repo) getAll(pageIndex, pageSize int) ([]*post, error) {
	query := `select p.id, p.title, p.content, u.id, u.username, u.email
			  from post p
         		left join user u on p.author_id = u.id
			  limit ? offset ?`
	rows, err := db.DB.Query(query, pageSize, pageIndex*pageSize)
	if err != nil {
		return nil, err
	}

	var posts []*post
	for rows.Next() {
		var p post
		err := rows.Scan(&p.id, &p.title, &p.content, &p.author.id, &p.author.username, &p.author.email)
		if err != nil {
			log.Println("Error scanning for post:", err)
			continue
		}
		posts = append(posts, &p)
	}

	// if there is an error yet there are some posts
	// it is better to return the posts and simply log the error
	// otherwise there has been a big error
	// so it is better to return the error
	if err = rows.Err(); err != nil {
		if posts != nil {
			log.Println("Error getting posts:", err)
			return posts, nil
		} else {
			return nil, err
		}
	}

	if posts == nil {
		posts = []*post{}
	}
	return posts, nil
}
