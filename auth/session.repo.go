package auth

import "forum/db"

type sessionRepo struct{}

func newSessionRepo() *sessionRepo {
	return &sessionRepo{}
}

func (r *sessionRepo) save(s *session) error {
	query := "insert into session (id, user_id, expires_at) values (?, ?, ?)"
	_, err := db.DB.Exec(query, s.id, s.userId, s.expiresAt)
	return err
}

func (r *sessionRepo) getById(id string) (*session, error) {
	query := "select id, user_id, expires_at from session where id = ?"
	row := db.DB.QueryRow(query, id)
	var s session
	err := row.Scan(&s.id, &s.userId, &s.expiresAt)
	return &s, err
}

func (r *sessionRepo) getByUserId(userId int) (*session, error) {
	query := "select id, user_id, expires_at from session where user_id = ?"
	row := db.DB.QueryRow(query, userId)
	var s session
	err := row.Scan(&s.id, &s.userId, &s.expiresAt)
	return &s, err
}
