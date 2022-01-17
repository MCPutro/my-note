package service

import (
	"context"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"time"
)

type NoteService struct{}

func (n NoteService) InsertNewNote(newNote entity.Note) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	newNote.Visible = true
	newNote.CreatedDate = time.Now()
	newNote.UpdateDate = time.Now()

	result, err := noteRepo.Insert(context.Background(), newNote)
	if err != nil {
		return nil, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n NoteService) UpdateNote(note entity.Note) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	note.UpdateDate = time.Now()

	result, err := noteRepo.Update(context.Background(), note)
	if err != nil {
		return nil, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n NoteService) GetNoteByUID(UserId string) (interface{}, error) {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	result, err := noteRepo.FindByUserId(context.Background(), UserId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (n NoteService) Remove(NoteId int) error {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	err := noteRepo.Remove(context.Background(), NoteId)

	if err != nil {
		return err
	}

	return nil
}

func (n NoteService) RemovePermanent(NoteId int) error {

	noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	err := noteRepo.RemovePermanent(context.Background(), NoteId)

	if err != nil {
		return err
	}

	return nil
}
