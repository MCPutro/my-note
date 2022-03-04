package repository

import (
	"context"
	"github.com/MCPutro/my-note/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, DB *gorm.DB, user entity.User) (*entity.User, error)
	FindById(ctx context.Context, DB *gorm.DB, id string) (*entity.User, error)
	FindByEmail(ctx context.Context, DB *gorm.DB, email string) (*entity.User, error)
	FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.User, error)
}
