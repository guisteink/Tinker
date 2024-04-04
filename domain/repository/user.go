package repository

import (
	"context"

	"github.com/guisteink/tinker/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	DeleteById(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*model.User, error)
	FindById() (*model.User, error)
}
