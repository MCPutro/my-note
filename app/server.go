package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func NewServer(DB *gorm.DB, router *mux.Router) *Server {
	return &Server{DB: DB, Router: router}
}
