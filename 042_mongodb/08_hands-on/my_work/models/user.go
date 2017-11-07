package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	log.Println("Trying to save:", m)
}

func LoadUsers() map[string]User {
	var jd Users
	data := make(map[string]User)

	f, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	// log.Println("Loaded:", string(f))
	json.Unmarshal(f, &jd)
	// log.Println("jd:", jd)
	for _, u := range jd {
		log.Println(u)
		data[u.ID] = u
	}
	log.Println("Loaded data:", data)
	return data
}
