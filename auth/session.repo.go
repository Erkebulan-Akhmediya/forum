package auth

import "forum/db"

type sessionRepo struct{}

func newSessionRepo() *sessionRepo {
	return &sessionRepo{}
}

func (r *sessionRepo) save(session session) error {
	query := "insert into session (id, user_id, expires_at) values (?, ?, ?)"
	_, err := db.DB.Exec(query, session.id, session.userId, session.expiresAt)
	return err
}
