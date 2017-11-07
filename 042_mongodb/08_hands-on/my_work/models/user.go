package models

import (
	"encoding/json"
	"io/ioutil"
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
	jm, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err.Error())
	}
	f.Write(jm)
}

func LoadUsers() map[string]User {
	var jd Users
	data := make(map[string]User)

	f, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	json.Unmarshal(f, &jd)
	for _, u := range jd {
		log.Println(u)
		data[u.ID] = u
	}
	log.Println("Loaded data:", data)
	return data
}
