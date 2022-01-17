package main

import (
	"fmt"
	"github.com/MCPutro/my-note/controller"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/schema"
	"github.com/MCPutro/my-note/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	myRoute = mux.NewRouter().StrictSlash(true)

	userService    = service.UserService{}
	userController = controller.UserController{
		Route:       myRoute,
		UserService: &userService,
	}

	noteService    = service.NoteService{}
	noteController = controller.NoteController{
		Route:       myRoute,
		NoteService: &noteService,
	}

	graphql = schema.GraphQL{
		UserService: &userService,
		NoteService: &noteService,
		Route:       myRoute,
	}
)

func init() {
	db := db_driver.GetConnection()
	defer db_driver.CloseConnection(db)
	db.AutoMigrate(entity.Note{})
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	userController.InitialPath("/user")

	noteController.InitialPath("/note")

	graphql.InitialPath("/graphql")

	myRoute.HandleFunc("/", chectAPI).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+PORT, myRoute))

}

func chectAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
