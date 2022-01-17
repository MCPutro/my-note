package service

import (
	"context"
	"errors"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
)

type UserService struct{}

func (uS UserService) CreateNewUser(newUser entity.User) (entity.User, error) {

	userRepo := repository.UserRepository(db_driver.GetConnection())

	ctx := context.Background()

	insert, err := userRepo.Insert(ctx, newUser)

	if err != nil {
		return entity.User{}, err
	}

	insert.Password = "*****"
	return insert, nil
}

func (uS UserService) SignInUser(user entity.User) (entity.User, error) {
	existingUser, err := getByEmail(user.Email)
	if err != nil {
		return entity.User{}, err
	}

	if existingUser.Password == user.Password {
		return existingUser, nil
	} else {
		return entity.User{}, errors.New("password salah")
	}
}

func getByEmail(email string) (entity.User, error) {
	userRepo := repository.UserRepository(db_driver.GetConnection())

	result, err := userRepo.FindByEmail(context.Background(), email)

	if err != nil {
		return entity.User{}, err
	}

	return result, nil

}

func (uS UserService) GetAllUser() ([]entity.User, error) {
	userRepo := repository.UserRepository(db_driver.GetConnection())

	result, err := userRepo.FindAll(context.Background())

	if err != nil {
		return []entity.User{}, nil
	}

	return result, nil

}
