package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
)

type NoteRepositoryInterface interface {
	Insert(ctx context.Context, newNote entity.Note) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Remove(ctx context.Context, noteId int) error
	RemovePermanent(ctx context.Context, noteId int) error
	FindByUserId(ctx context.Context, userId string) ([]entity.Note, error)
}
