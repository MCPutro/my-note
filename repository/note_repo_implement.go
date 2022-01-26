package repository

import (
	"context"
	"fmt"
	"github.com/MCPutro/my-note/entity"
	"gorm.io/gorm"
)

type noteRepoImplement struct {
	DB *gorm.DB
}

func (n *noteRepoImplement) Insert(ctx context.Context, newNote entity.Note) (entity.Note, error) {

	result := n.DB.WithContext(ctx).Create(&newNote)

	defer func() {
		db, _ := n.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if result.Error != nil {
		return entity.Note{}, result.Error
	}

	return entity.Note{
		ID:          newNote.ID,
		Text:        newNote.Text,
		Visible:     newNote.Visible,
		CreatedDate: newNote.CreatedDate,
		UpdateDate:  newNote.UpdateDate,
	}, nil
}

func (n *noteRepoImplement) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	result := n.DB.WithContext(ctx).Updates(&note)

	defer func() {
		db, _ := n.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if result.Error != nil {
		return entity.Note{}, result.Error
	}

	return entity.Note{
		ID:          note.ID,
		Text:        note.Text,
		Visible:     note.Visible,
		CreatedDate: note.CreatedDate,
		UpdateDate:  note.UpdateDate,
	}, nil
}

func (n *noteRepoImplement) Remove(ctx context.Context, noteId int) error { //set visible to false

	var note entity.Note

	result := n.DB.WithContext(ctx).Where("ID = ?", noteId).First(&note).Update("Visible", false)

	defer func() {
		db, _ := n.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (n *noteRepoImplement) RemovePermanent(ctx context.Context, noteId int) error {
	result := n.DB.WithContext(ctx).Where("ID = ? AND Visible = ?", noteId, false).Delete(&entity.Note{})

	defer func() {
		db, _ := n.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (n *noteRepoImplement) FindByUserId(ctx context.Context, userId string) ([]entity.Note, error) {

	var listNote []entity.Note
	find := n.DB.WithContext(ctx).Where(entity.Note{UserId: userId}).Find(&listNote)
	defer func() {
		db, _ := n.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if find.Error != nil {
		return listNote, find.Error
	}

	return listNote, nil
}

func GetNoteRepository(db *gorm.DB) NoteRepositoryInterface {
	return &noteRepoImplement{DB: db}
}
