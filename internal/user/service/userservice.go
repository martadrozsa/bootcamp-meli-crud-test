package service

import (
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
)

type serviceImpl struct {
	repository domain.UserRepository
}

func CreateUserService(r domain.UserRepository) domain.UserService {
	return &serviceImpl{r}
}

func (s *serviceImpl) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *serviceImpl) GetById(id int64) (*domain.User, error) {
	userId, err := s.repository.GetById(id)

	if err != nil {
		return nil, err
	}
	return userId, nil
}

func (s *serviceImpl) Create(name string, age int, movieGenre string) (*domain.User, error) {
	newUser, err := s.repository.Create(name, age, movieGenre)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *serviceImpl) UpdateAge(id int64, age int) (*domain.User, error) {
	return s.repository.UpdateAge(id, age)
}

func (s *serviceImpl) Delete(id int64) error {
	return s.repository.Delete(id)
}
