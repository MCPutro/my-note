package main

import (
	"fmt"
	"github.com/MCPutro/my-note/controller"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"github.com/MCPutro/my-note/schema"
	"github.com/MCPutro/my-note/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	myRoute = mux.NewRouter()

	db = db_driver.GetConnection()

	userRepo       = repository.NewUserRepository()
	userService    = service.NewUserService(userRepo, db)
	userController = controller.NewUserController(myRoute, userService)

	noteRepo       = repository.NewNoteRepository()
	noteService    = service.NewNoteService(noteRepo, db)
	noteController = controller.NewNoteController(myRoute, noteService)
)

func init() {
	err := db.AutoMigrate(entity.Note{})
	if err != nil {
		log.Fatal("error init main : ", err)
		return
	}
}

func NewGraphQL(userService service.UserService, noteService service.NoteService, route *mux.Router) *schema.GraphQL {
	return &schema.GraphQL{
		UserService: userService,
		NoteService: noteService,
		Route:       route,
	}
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	userController.InitialPath("/user")
	noteController.InitialPath("/note")

	graphql := NewGraphQL(userService, noteService, myRoute)
	graphql.InitialPath("/graphql")

	myRoute.HandleFunc("/", checkAPI).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+PORT, myRoute))

}

func checkAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
