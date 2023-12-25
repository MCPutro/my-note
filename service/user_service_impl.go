package service

import (
	"context"
	"errors"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/repository"
	"github.com/MCPutro/my-note/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB}
}

func (us *UserServiceImpl) CreateNewUser(ctx context.Context, newUser entity.User) (*entity.User, error) {

	//userRepo := repository.UserRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	//Encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)

	insert, err := us.UserRepository.Save(ctx2, us.DB, newUser)

	if err != nil {
		return nil, err
	}

	insert.Password = "*****"
	return insert, nil
}

func (us *UserServiceImpl) SignInUser(ctx context.Context, userLLogin entity.User) (*entity.User, error) {

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	existingUser, err := us.getByEmail(ctx2, userLLogin.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userLLogin.Password)); err == nil {
		// existingUser.Password == user.Password {
		token, err := util.GenerateToken(existingUser.ID)
		if err != nil {
			return nil, err
		}
		existingUser.Token = token
		return existingUser, nil
	} else {
		return nil, errors.New("invalid  email or password")
	}
}

func (us *UserServiceImpl) getByEmail(ctx context.Context, email string) (*entity.User, error) {
	//userRepo := repository.UserRepository(db_driver.GetConnection())

	result, err := us.UserRepository.FindByEmail(ctx, us.DB, email)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (us *UserServiceImpl) GetAllUser(ctx context.Context) (*[]entity.User, error) {
	//userRepo := repository.UserRepository(db_driver.GetConnection())

	ctx2, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	result, err := us.UserRepository.FindAll(ctx2, us.DB)

	if err != nil {
		return nil, err
	}

	return result, nil

}
