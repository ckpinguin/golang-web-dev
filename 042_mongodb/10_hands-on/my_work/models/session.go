package models

import "time"

type session struct {
	un           string
	lastActivity time.Time
}
