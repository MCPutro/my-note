package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type UserController interface {
	InitialPath(route *mux.Router, path string)
	createNewUser(w http.ResponseWriter, r *http.Request)
	signInUser(w http.ResponseWriter, r *http.Request)
	getAllUser(w http.ResponseWriter, r *http.Request)
}
