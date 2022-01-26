package service

import (
	"context"
	"errors"
	db_driver "github.com/MCPutro/my-note/db-driver"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"time"
)

type UserService struct {
	CtxParent context.Context
}

func (uS UserService) CreateNewUser(newUser entity.User) (entity.User, error) {

	userRepo := repository.UserRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(uS.CtxParent, 10*time.Second)
	defer cancelFunc()

	insert, err := userRepo.Insert(ctx, newUser)

	if err != nil {
		return entity.User{}, err
	}

	insert.Password = "*****"
	return insert, nil
}

func (uS UserService) SignInUser(user entity.User) (entity.User, error) {

	ctx, cancelFunc := context.WithTimeout(uS.CtxParent, 10*time.Second)
	defer cancelFunc()

	existingUser, err := getByEmail(ctx, user.Email)
	if err != nil {
		return entity.User{}, err
	}

	if existingUser.Password == user.Password {
		return existingUser, nil
	} else {
		return entity.User{}, errors.New("password salah")
	}
}

func getByEmail(ctx context.Context, email string) (entity.User, error) {
	userRepo := repository.UserRepository(db_driver.GetConnection())

	result, err := userRepo.FindByEmail(ctx, email)

	if err != nil {
		return entity.User{}, err
	}

	return result, nil

}

func (uS UserService) GetAllUser() ([]entity.User, error) {
	userRepo := repository.UserRepository(db_driver.GetConnection())

	ctx, cancelFunc := context.WithTimeout(uS.CtxParent, 10*time.Second)
	defer cancelFunc()

	result, err := userRepo.FindAll(ctx)

	if err != nil {
		return []entity.User{}, err
	}

	return result, nil

}
