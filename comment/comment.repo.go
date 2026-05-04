package comment

import (
	"forum/db"
	"forum/utils"
	"log"
)

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) savePostComment(c *postComment) error {
	query := `insert into comment (content, author_id, post_id) 
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, c.content, c.authorId, c.postId)
	return row.Scan(&c.id)
}

func (r *repo) saveReplyComment(c *replyComment) error {
	query := `insert into comment (content, author_id, comment_id)
			  values (?, ?, ?)
			  returning id`
	row := db.DB.QueryRow(query, c.content, c.authorId, c.commentId)
	return row.Scan(&c.id)
}

func (r *repo) getAllByPostId(postId int, page *utils.Page) ([]*postComment, error) {
	query := `select c.id, c.content, u.id, u.username, u.email, f.id
			  from comment c
         		left join user u on c.author_id = u.id
         		left join file f on c.id = f.comment_id
			  where c.post_id = ?
			  limit ? offset ?`

	rows, err := db.DB.Query(query, postId, page.Size, page.Index*page.Size)
	if err != nil {
		return nil, err
	}

	var comments []*postComment
	for rows.Next() {
		var c postComment
		err := rows.Scan(
			&c.id,
			&c.content,
			&c.author.id,
			&c.author.username,
			&c.author.email,
			&c.fileId,
		)
		if err != nil {
			log.Println("Error scanning for comment:", err)
			continue
		}
		comments = append(comments, &c)
	}

	// if there is an error yet there are some comments
	// it is better to return the comments and simply log the error
	// otherwise there has been a big error
	// so it is better to return the error
	if err = rows.Err(); err != nil {
		if comments != nil {
			log.Println("Error getting posts:", err)
			return comments, nil
		} else {
			return nil, err
		}
	}

	if comments == nil {
		comments = []*postComment{}
	}
	return comments, nil
}
