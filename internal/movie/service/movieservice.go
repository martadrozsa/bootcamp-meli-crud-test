package service

import (
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
)

type serviceImpl struct {
	repository domain.MovieRepository
}

func CreateMovieService(r domain.MovieRepository) domain.MovieService {
	return &serviceImpl{r}
}

func (s *serviceImpl) GetAll() ([]*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) GetById(id int64) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Create(name string, genre string, year int, award string) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) UpdateAward(id int64, award string) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
