package model

import "time"

type (
	LoginSession struct {
		Id           string
		UserId       string
		RefreshToken string
		Revoked      bool
		ExpiresAt    time.Time
	}
)
