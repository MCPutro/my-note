package controller

import (
	"net/http"
)

type UserController interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	SignInUser(w http.ResponseWriter, r *http.Request)
	GetAllUser(w http.ResponseWriter, r *http.Request)
}
