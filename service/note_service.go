package service

import (
	"context"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"time"
)

type NoteService struct {
	CtxParent context.Context
}

func (n NoteService) InsertNewNote(newNote entity.Note) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(n.CtxParent, 10*time.Second)
	defer cancelFunc()

	newNote.Visible = true
	newNote.CreatedDate = time.Now()
	newNote.UpdateDate = time.Now()

	result, err := noteRepo.Insert(ctx, newNote)
	if err != nil {
		return nil, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n NoteService) UpdateNote(note entity.Note) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(n.CtxParent, 10*time.Second)
	defer cancelFunc()

	note.UpdateDate = time.Now()

	result, err := noteRepo.Update(ctx, note)
	if err != nil {
		return nil, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n NoteService) GetNoteByUID(UserId string) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(n.CtxParent, 10*time.Second)
	defer cancelFunc()

	result, err := noteRepo.FindByUserId(ctx, UserId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (n NoteService) Remove(NoteId int) error {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(n.CtxParent, 10*time.Second)
	defer cancelFunc()

	err := noteRepo.Remove(ctx, NoteId)

	if err != nil {
		return err
	}

	return nil
}

func (n NoteService) RemovePermanent(NoteId int) error {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(n.CtxParent, 10*time.Second)
	defer cancelFunc()

	err := noteRepo.RemovePermanent(ctx, NoteId)

	if err != nil {
		return err
	}

	return nil
}
