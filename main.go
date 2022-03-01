package main

import (
	"fmt"
	"github.com/MCPutro/my-note/app"
	"github.com/MCPutro/my-note/controller"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"github.com/MCPutro/my-note/service"
	"log"
	"net/http"
	"os"
)

func main2() {

	db := db_driver.GetConnection()

	err := db.AutoMigrate(entity.Note{})

	if err != nil {
	}

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, db)
	userController := controller.NewUserController(userService)

	noteRepo := repository.NewNoteRepository()
	noteService := service.NewNoteService(noteRepo, db)
	noteController := controller.NewNoteController(noteService)

	graphql := app.NewGraphQL(userService, noteService)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	server := app.NewRouter(userController, noteController, graphql)

	log.Fatal(http.ListenAndServe(":"+PORT, server))

}

func main() {
	server := InitServer()
	err := server.DB.AutoMigrate(entity.Note{})
	if err != nil {
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, server.Router))
}
