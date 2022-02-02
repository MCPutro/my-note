package service

import (
	"context"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"gorm.io/gorm"
	"time"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *gorm.DB
}

func NewNoteService(noteRepository repository.NoteRepository, DB *gorm.DB) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB:             DB,
	}
}

func (n *NoteServiceImpl) InsertNewNote(ctx context.Context, newNote entity.Note) (entity.Note, error) {

	//noteRepo := repository.NewNoteRepoImplement(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	newNote.Visible = true
	newNote.CreatedDate = time.Now()
	newNote.UpdateDate = time.Now()

	result, err := n.NoteRepository.Save(ctx2, n.DB, newNote)
	if err != nil {
		return result, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n *NoteServiceImpl) UpdateNote(ctx context.Context, note entity.Note) (entity.Note, error) {

	//noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	note.UpdateDate = time.Now()

	result, err := n.NoteRepository.Update(ctx2, n.DB, note)
	if err != nil {
		return result, err
	}

	//fmt.Println("new note >>>>> ", result)

	return result, nil
}

func (n *NoteServiceImpl) GetNoteByUID(ctx context.Context, UserId string) ([]entity.Note, error) {

	//noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	result, err := n.NoteRepository.FindByUID(ctx2, n.DB, UserId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (n *NoteServiceImpl) Remove(ctx context.Context, NoteId int) error {

	//noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	err := n.NoteRepository.Delete(ctx2, n.DB, NoteId)

	if err != nil {
		return err
	}

	return nil
}

func (n *NoteServiceImpl) RemovePermanent(ctx context.Context, NoteId int) error {

	//noteRepo := repository.GetNoteRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	err := n.NoteRepository.DeletePermanent(ctx2, n.DB, NoteId)

	if err != nil {
		return err
	}

	return nil
}
