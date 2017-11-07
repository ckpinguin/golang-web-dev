package main

import (
	"net/http"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/06_hands-on/my_work/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
