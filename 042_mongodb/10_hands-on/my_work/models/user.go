package models

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}
