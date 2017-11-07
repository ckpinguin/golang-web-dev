package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/08_hands-on/my_work/controllers"
	"github.com/ckpinguin/golang-web-dev/042_mongodb/08_hands-on/my_work/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	data := make(map[string]models.User)

	f, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Loaded:", string(f))
	json.Unmarshal(f, &data)
	log.Println("data:", data)
	return data
}
