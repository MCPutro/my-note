package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"io"
	"net/http"
	"strconv"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{NoteService: noteService}
}

func (nc *NoteControllerImpl) CreateNew(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := io.ReadAll(r.Body)
	var newNote entity.Note
	json.Unmarshal(requestPayload, &newNote)

	note, err := nc.NoteService.InsertNewNote(r.Context(), newNote)
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

func (nc *NoteControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := io.ReadAll(r.Body)
	var note entity.Note
	json.Unmarshal(requestPayload, &note)

	updateNote, err := nc.NoteService.UpdateNote(r.Context(), note)
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

func (nc *NoteControllerImpl) GetByUserId(w http.ResponseWriter, r *http.Request) {
	UserId := r.Header.Get("UserId")

	noteByUID, err := nc.NoteService.GetNoteByUID(r.Context(), UserId)
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

func (nc *NoteControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(r.Header.Get("NoteId"))
	if err != nil {
		fmt.Fprint(w, "harus bilangan bulat")
		return
	}
	err = nc.NoteService.Remove(r.Context(), noteId)
	if err != nil {
		fmt.Fprint(w, "error nih : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (nc *NoteControllerImpl) DeletePermanent(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(r.Header.Get("NoteId"))
	if err != nil {
		fmt.Fprint(w, "harus bilangan bulat")
		return
	}
	err = nc.NoteService.RemovePermanent(r.Context(), noteId)
	if err != nil {
		fmt.Fprint(w, "error nih : ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

//func test(w http.ResponseWriter, r *http.Request) {
//	param := mux.Vars(r)
//	sm := param["sm"]
//	fmt.Println(sm)
//}
