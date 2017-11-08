package sessions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ckpinguin/golang-web-dev/042_mongodb/02_json/models"
	"github.com/satori/go.uuid"
)

const Length int = 30

var Users = map[string]models.User{}       // user ID, user
var Sessions = map[string]models.Session{} // session ID, session
var LastCleaned time.Time

func getUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = Length
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.lastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.un]
	// refresh session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	SessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
