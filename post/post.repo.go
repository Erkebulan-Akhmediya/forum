package post

import (
	"encoding/json"
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

func (r *repo) existsById(id int) (bool, error) {
	query := "select exists(select id from post where id = ?)"
	row := db.DB.QueryRow(query, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *repo) getAll(pageIndex, pageSize int) ([]*post, error) {
	query := `select 
				p.id, 
				p.title, 
				p.content, 
				u.id, 
				u.username, 
				u.email, 
				json_group_array(f.id)
			  from post p
         		left join user u on p.author_id = u.id
         		left join file f on p.id = f.post_id
			  group by p.id
			  limit ? offset ?`
	rows, err := db.DB.Query(query, pageSize, pageIndex*pageSize)
	if err != nil {
		return nil, err
	}

	var posts []*post
	for rows.Next() {
		var p post

		// since sqlite returns json array as string
		// the array is written to string variable (rawFileIds)
		var rawFileIds string

		err := rows.Scan(
			&p.id,
			&p.title,
			&p.content,
			&p.author.id,
			&p.author.username,
			&p.author.email,
			&rawFileIds,
		)
		if err != nil {
			log.Println("Error scanning for post:", err)
			continue
		}

		// since it is better to represent file ids as int slice
		// the string varialbe is parsed into int slice (fileIds)
		var fileIds []int
		if err := json.Unmarshal([]byte(rawFileIds), &fileIds); err != nil {
			log.Println("Error parsing file id array", err)
			continue
		}

		// since sqlite returns an empty json array as an array containing one null value
		// which then parsed into a slice containing single 0 value
		// that slice is changed to an empty slice
		if len(fileIds) == 1 && fileIds[0] == 0 {
			fileIds = []int{}
		}
		p.fileIds = fileIds
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
