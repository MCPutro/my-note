package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type NoteController interface {
	InitialPath(route *mux.Router, path string)
	createNewNote(w http.ResponseWriter, r *http.Request)
	updateNote(w http.ResponseWriter, r *http.Request)
	getNoteByUserId(w http.ResponseWriter, r *http.Request)
	remove(w http.ResponseWriter, r *http.Request)
	removePermanent(w http.ResponseWriter, r *http.Request)
}
