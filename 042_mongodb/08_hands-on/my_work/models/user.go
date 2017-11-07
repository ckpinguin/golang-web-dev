package models

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Age    int    `json:"Age"`
}

type Users []User

// Id was of type string before

func StoreUsers(m map[string]User) {
	f, err := os.Create("./users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()
	log.Println("Saving:", m)
	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	data := make(map[string]User)

	f, err := os.Open("./users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&data)
	log.Println("Loaded data:", data)
	return data
}
