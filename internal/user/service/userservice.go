package service

import (
	"context"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
)

type serviceImpl struct {
	repository domain.UserRepository
}

func (s serviceImpl) GetAll(ctx context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) GetById(ctx context.Context, id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) Create(ctx context.Context, name string, age int, movieGenre string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) UpdateAge(ctx context.Context, id int64, age int) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImpl) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func CreateUserService(r domain.UserRepository) domain.UserService {
	return &serviceImpl{r}
}
