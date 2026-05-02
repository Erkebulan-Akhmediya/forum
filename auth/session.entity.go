package auth

import "time"

type session struct {
	id        string
	userId    int
	expiresAt time.Time
}
