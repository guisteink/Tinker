package service

import (
	"context"

	"github.com/guisteink/tinker/domain/model"
	"github.com/guisteink/tinker/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func newUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteById(ctx, id)
}

func (s *UserService) FindAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) FindUserById(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.FindById(ctx, id)
}
