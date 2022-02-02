package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
	"gorm.io/gorm"
)

type noteRepoImplement struct {
}

func NewNoteRepository() NoteRepository {
	return &noteRepoImplement{}
}

func (n *noteRepoImplement) Save(ctx context.Context, DB *gorm.DB, newNote entity.Note) (entity.Note, error) {

	result := DB.WithContext(ctx).Create(&newNote)

	//defer func() {
	//	db, _ := n.DB.DB()
	//	db.Close()
	//	fmt.Println("Close connection to db")
	//}()

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

func (n *noteRepoImplement) Update(ctx context.Context, DB *gorm.DB, note entity.Note) (entity.Note, error) {
	result := DB.WithContext(ctx).Updates(&note)

	//defer func() {
	//	db, _ := n.DB.DB()
	//	db.Close()
	//	fmt.Println("Close connection to db")
	//}()

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

func (n *noteRepoImplement) Delete(ctx context.Context, DB *gorm.DB, noteId int) error { //set visible to false

	var note entity.Note

	result := DB.WithContext(ctx).Where("ID = ?", noteId).First(&note).Update("Visible", false)

	//defer func() {
	//	db, _ := n.DB.DB()
	//	db.Close()
	//	fmt.Println("Close connection to db")
	//}()

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (n *noteRepoImplement) DeletePermanent(ctx context.Context, DB *gorm.DB, noteId int) error {
	result := DB.WithContext(ctx).Where("ID = ? AND Visible = ?", noteId, false).Delete(&entity.Note{})

	//defer func() {
	//	db, _ := n.DB.DB()
	//	db.Close()
	//	fmt.Println("Close connection to db")
	//}()

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (n *noteRepoImplement) FindByUID(ctx context.Context, DB *gorm.DB, userId string) ([]entity.Note, error) {

	var listNote []entity.Note
	find := DB.WithContext(ctx).Where(entity.Note{UserId: userId}).Find(&listNote)
	//defer func() {
	//	db, _ := n.DB.DB()
	//	db.Close()
	//	fmt.Println("Close connection to db")
	//}()

	if find.Error != nil {
		return listNote, find.Error
	}

	return listNote, nil
}
