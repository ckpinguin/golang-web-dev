package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/06_hands-on/my_work/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
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

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
