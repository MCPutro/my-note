package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
)

type UserRepositoryInterface interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	FindById(ctx context.Context, id string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
}
