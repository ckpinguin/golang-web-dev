package main

import (
	"net/http"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/02_json/models"
	"github.com/ckpinguin/golang-web-dev/042_mongodb/06_hands-on/my_work/controllers"
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
	return make(map[string]models.User)
}
