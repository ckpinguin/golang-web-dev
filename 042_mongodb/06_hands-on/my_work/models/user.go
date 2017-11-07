package models

// import "gopkg.in/mgo.v2/bson"

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// Id was of type string before
