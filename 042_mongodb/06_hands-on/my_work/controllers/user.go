package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/06_hands-on/my_work/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	if u, ok := uc.session[id]; ok {
		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", uj)
		return
	}
	w.WriteHeader(http.StatusNotFound) // 404
	return
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.ID = uuid.NewV4().String()

	// store the user in mongodb
	uc.session[u.ID] = u

	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if _, ok := uc.session[id]; ok {
		// Delete user
		delete(uc.session, id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprint(w, "Deleted user", id, "\n")
		return
	}
	w.WriteHeader(http.StatusNotFound) // 404
	return
}
