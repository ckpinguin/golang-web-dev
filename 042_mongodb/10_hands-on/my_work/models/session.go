package models


type session struct {
	un           string
	lastActivity time.Time
}