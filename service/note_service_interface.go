package service

import (
	"context"
	"github.com/MCPutro/my-note/entity"
)

type NoteService interface {
	InsertNewNote(ctx context.Context, newNote entity.Note) (*entity.Note, error)
	UpdateNote(ctx context.Context, note entity.Note) (*entity.Note, error)
	GetNoteByUID(ctx context.Context, UserId string) (*[]entity.Note, error)
	Remove(ctx context.Context, NoteId int) error
	RemovePermanent(ctx context.Context, NoteId int) error
}
