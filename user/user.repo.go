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

func (r *repo) existsByEmail(email string) (bool, error) {
	query := "select exists(select id from user where email = ?)"
	row := db.DB.QueryRow(query, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *repo) getByEmail(email string) (*User, error) {
	query := "select id, username, password, email from user where email = ?"
	row := db.DB.QueryRow(query, email)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return &user, err
}
