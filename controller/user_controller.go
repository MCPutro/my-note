package controller

import (
	"encoding/json"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	Route       *mux.Router
	UserService *service.UserService
}

func (uc *UserController) InitialPath(path string) {
	uc.Route.HandleFunc(path+"/signup", uc.createNewUser).Methods("POST")
	uc.Route.HandleFunc(path+"/signin", uc.signInUser).Methods("POST")
	uc.Route.HandleFunc(path+"/getAllUser", uc.getAllUser).Methods("GET")
}

func (uc UserController) createNewUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	var newUser entity.User
	json.Unmarshal(requestPayload, &newUser)

	result, err := uc.UserService.CreateNewUser(newUser)

	var respJson []byte

	if err != nil {
		//fmt.Fprint(w, err)
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "success",
				Data:   result,
			})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func (uc UserController) signInUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestPayload, &user)

	var respJson []byte
	inUser, err := uc.UserService.SignInUser(user)
	if err != nil {
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "error",
				Data:   inUser,
			})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)

}

func (uc UserController) getAllUser(w http.ResponseWriter, r *http.Request) {

	var respJson []byte
	listUser, err := uc.UserService.GetAllUser()
	if err != nil {
		respJson, _ = json.Marshal(
			response.Resp{
				Status:  "error",
				Message: err.Error(),
			})
	} else {
		respJson, _ = json.Marshal(
			response.Resp{
				Status: "error",
				Data:   listUser,
			})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("oke", "haha") //set data to header resp
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)

}
