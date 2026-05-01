package user

import "forum/db"

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) save(user *User) error {
	query := `insert into user (username, password, email)
			  values (?, ?, ?)`

	_, err := db.DB.Exec(query, user.Username, user.Password, user.Email)
	return err
}
