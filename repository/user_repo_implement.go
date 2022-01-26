package repository

import (
	"context"
	"fmt"
	"github.com/MCPutro/my-note/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepoImplement struct {
	//UserIdentity entity.User
	DB *gorm.DB
}

func (repo *userRepoImplement) Insert(ctx context.Context, user entity.User) (entity.User, error) {

	newUUID := uuid.New().String()
	user.ID = newUUID

	create := repo.DB.WithContext(ctx).Create(&user)

	defer func() {
		db, _ := repo.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if create.Error != nil {
		return entity.User{}, create.Error
	}

	return entity.User{ID: user.ID, Email: user.Email}, nil
}

func (repo *userRepoImplement) FindById(ctx context.Context, id string) (entity.User, error) {
	existingUser := entity.User{}

	firstResult := repo.DB.WithContext(ctx).Where("id = ?", id).First(&existingUser)

	defer func() {
		db, _ := repo.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if firstResult.Error != nil {
		return entity.User{}, firstResult.Error
	}

	return existingUser, nil
}

func (repo *userRepoImplement) FindByEmail(ctx context.Context, email string) (entity.User, error) {

	existingUser := entity.User{}

	firstResult := repo.DB.WithContext(ctx).Where("email = ?", email).First(&existingUser)

	defer func() {
		db, _ := repo.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if firstResult.Error != nil {
		return entity.User{}, firstResult.Error
	}

	return existingUser, nil
}

func (repo *userRepoImplement) FindAll(ctx context.Context) ([]entity.User, error) {

	var listUser []entity.User

	find := repo.DB.WithContext(ctx).Find(&listUser)

	defer func() {
		db, _ := repo.DB.DB()
		db.Close()
		fmt.Println("Close connection to db")
	}()

	if find.Error != nil {
		return listUser, find.Error
	}

	return listUser, nil

}

func UserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepoImplement{DB: db}
}
