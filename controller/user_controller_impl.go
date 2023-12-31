package controller

import (
	"encoding/json"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"io"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (u *UserControllerImpl) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := io.ReadAll(r.Body)
	var newUser entity.User
	json.Unmarshal(requestPayload, &newUser)

	result, err := u.UserService.CreateNewUser(r.Context(), newUser)

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

func (u *UserControllerImpl) SignInUser(w http.ResponseWriter, r *http.Request) {
	requestPayload, _ := io.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestPayload, &user)

	var respJson []byte
	inUser, err := u.UserService.SignInUser(r.Context(), user)
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

func (u *UserControllerImpl) GetAllUser(w http.ResponseWriter, r *http.Request) {
	var respJson []byte
	listUser, err := u.UserService.GetAllUser(r.Context())
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
