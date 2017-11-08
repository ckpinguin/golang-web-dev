package models

import "time"

type Session struct {
	un           string
	lastActivity time.Time
}
