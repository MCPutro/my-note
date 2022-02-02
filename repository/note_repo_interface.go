package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(ctx context.Context, DB *gorm.DB, newNote entity.Note) (entity.Note, error)
	Update(ctx context.Context, DB *gorm.DB, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, DB *gorm.DB, noteId int) error
	DeletePermanent(ctx context.Context, DB *gorm.DB, noteId int) error
	FindByUID(ctx context.Context, DB *gorm.DB, userId string) ([]entity.Note, error)
}
