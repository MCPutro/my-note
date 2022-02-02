package app

import (
	"fmt"
	"github.com/MCPutro/my-note/controller"
	"github.com/MCPutro/my-note/schema"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userController controller.UserController, noteController controller.NoteController, graphql *schema.GraphQL) *mux.Router {
	myRoute := mux.NewRouter()

	userController.InitialPath(myRoute, "/user")
	noteController.InitialPath(myRoute, "/note")
	graphql.InitialPath(myRoute, "/graphql")

	myRoute.HandleFunc("/", checkAPI)
	return myRoute
}

func checkAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
