package service

import (
	"context"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
)

type serviceImpl struct {
	repository domain.MovieRepository
}

func CreateMovieService(r domain.MovieRepository) domain.MovieService {
	return &serviceImpl{
		repository: r}
}

func (s serviceImpl) GetAll(ctx context.Context) (*[]domain.Movie, error) {
	movies, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s serviceImpl) GetById(ctx context.Context, id int64) (*domain.Movie, error) {
	movie, err := s.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s serviceImpl) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	newMovie, err := s.repository.Create(ctx, movie)

	if err != nil {
		return nil, err
	}

	return newMovie, nil
}

func (s serviceImpl) UpdateAward(ctx context.Context, id int64, award int) (*domain.Movie, error) {

	movie := domain.Movie{
		Id:    id,
		Award: award,
	}

	movieUpdate, err := s.repository.UpdateAward(ctx, &movie)

	if err != nil {
		return nil, err
	}

	return movieUpdate, nil
}

func (s serviceImpl) Delete(ctx context.Context, id int64) error {

	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
