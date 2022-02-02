package controller

import "net/http"

type NoteController interface {
	InitialPath(path string)
	createNewNote(w http.ResponseWriter, r *http.Request)
	updateNote(w http.ResponseWriter, r *http.Request)
	getNoteByUserId(w http.ResponseWriter, r *http.Request)
	remove(w http.ResponseWriter, r *http.Request)
	removePermanent(w http.ResponseWriter, r *http.Request)
}
