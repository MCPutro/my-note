package controller

import (
	"net/http"
)

type NoteController interface {
	CreateNew(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeletePermanent(w http.ResponseWriter, r *http.Request)
}
