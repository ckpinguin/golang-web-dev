package models

type User struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Age    int    `json:"Age"`
}

// Id was of type string before
