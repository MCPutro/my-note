package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepoImplement struct {
}

func NewUserRepository() UserRepository {
	return &userRepoImplement{}
}

func (repo *userRepoImplement) Save(ctx context.Context, DB *gorm.DB, user entity.User) (entity.User, error) {

	newUUID := uuid.New().String()
	user.ID = newUUID

	create := DB.WithContext(ctx).Create(&user)

	if create.Error != nil {
		return entity.User{}, create.Error
	}

	return entity.User{ID: user.ID, Email: user.Email}, nil
}

func (repo *userRepoImplement) FindById(ctx context.Context, DB *gorm.DB, id string) (entity.User, error) {
	existingUser := entity.User{}

	firstResult := DB.WithContext(ctx).Where("id = ?", id).First(&existingUser)

	if firstResult.Error != nil {
		return entity.User{}, firstResult.Error
	}

	return existingUser, nil
}

func (repo *userRepoImplement) FindByEmail(ctx context.Context, DB *gorm.DB, email string) (entity.User, error) {

	existingUser := entity.User{}

	firstResult := DB.WithContext(ctx).Where("email = ?", email).First(&existingUser)

	if firstResult.Error != nil {
		return entity.User{}, firstResult.Error
	}

	return existingUser, nil
}

func (repo *userRepoImplement) FindAll(ctx context.Context, DB *gorm.DB) ([]entity.User, error) {

	var listUser []entity.User

	find := DB.WithContext(ctx).Find(&listUser)

	if find.Error != nil {
		return listUser, find.Error
	}

	return listUser, nil

}
