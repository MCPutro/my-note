//go:build wireinject
// +build wireinject

package main

import (
	"github.com/MCPutro/my-note/app"
	"github.com/MCPutro/my-note/controller"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/repository"
	"github.com/MCPutro/my-note/service"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

var NoteSet = wire.NewSet(
	repository.NewNoteRepository,
	service.NewNoteService,
	controller.NewNoteController,
)

func InitServer() *app.Server {
	wire.Build(
		db_driver.GetConnection,
		UserSet,
		NoteSet,
		NewGraphQL,
		app.NewRouter,
		app.NewServer,
	)

	return nil

}
