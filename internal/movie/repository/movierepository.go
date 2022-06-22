package repository

import (
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
)

type repositoryImpl struct {
}

func CreateMovieRepository() domain.MovieRepository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll() ([]*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) GetById(id int64) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Create(name string, genre string, year int, award string) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) UpdateAward(id int64, award string) (*domain.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
