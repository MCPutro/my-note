package service

import (
	"context"
	"github.com/MCPutro/my-note/entity"
)

type UserService interface {
	CreateNewUser(ctx context.Context, newUser entity.User) (entity.User, error)
	SignInUser(ctx context.Context, user entity.User) (entity.User, error)
	getByEmail(ctx context.Context, email string) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
}
