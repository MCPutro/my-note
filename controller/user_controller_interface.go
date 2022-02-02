package controller

import "net/http"

type UserController interface {
	InitialPath(path string)
	createNewUser(w http.ResponseWriter, r *http.Request)
	signInUser(w http.ResponseWriter, r *http.Request)
	getAllUser(w http.ResponseWriter, r *http.Request)
}
