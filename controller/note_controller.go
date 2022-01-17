package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type NoteController struct {
	Route       *mux.Router
	NoteService *service.NoteService
}

func (nc *NoteController) InitialPath(path string) {
	nc.Route.HandleFunc(path+"/create", nc.createNewNote).Methods("POST")
	nc.Route.HandleFunc(path+"/update", nc.updateNote).Methods("POST")
	nc.Route.HandleFunc(path+"/getAllByUID", nc.getNoteByUserId).Methods("GET")
	nc.Route.HandleFunc(path+"/remove", nc.remove).Methods("GET")
	nc.Route.HandleFunc(path+"/removePermanent", nc.removePermanent).Methods("GET")
}

func (nc *NoteController) createNewNote(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	var newNote entity.Note
	json.Unmarshal(requestPayload, &newNote)

	note, err := nc.NoteService.InsertNewNote(newNote)
	var respJson []byte
	if err != nil {
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		//fmt.Println("after create >>> ", note)
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "success",
				Data:   note,
			})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (nc *NoteController) updateNote(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	var note entity.Note
	json.Unmarshal(requestPayload, &note)

	updateNote, err := nc.NoteService.UpdateNote(note)
	var respJson []byte
	if err != nil {
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		//fmt.Println("after create >>> ", note)
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "success",
				Data:   updateNote,
			})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (nc *NoteController) getNoteByUserId(w http.ResponseWriter, r *http.Request) {
	UserId := r.Header.Get("UserId")

	noteByUID, err := nc.NoteService.GetNoteByUID(UserId)
	var respJson []byte
	if err != nil {
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "success",
				Data:   noteByUID,
			})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (nc *NoteController) remove(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(r.Header.Get("NoteId"))
	if err != nil {
		fmt.Fprint(w, "harus bilangan bulat")
		return
	}
	err = nc.NoteService.Remove(noteId)
	if err != nil {
		fmt.Fprint(w, "error nih : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (nc *NoteController) removePermanent(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(r.Header.Get("NoteId"))
	if err != nil {
		fmt.Fprint(w, "harus bilangan bulat")
		return
	}
	err = nc.NoteService.RemovePermanent(noteId)
	if err != nil {
		fmt.Fprint(w, "error nih : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
