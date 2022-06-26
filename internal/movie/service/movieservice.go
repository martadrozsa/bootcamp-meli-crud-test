package service

import (
	"context"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
)

type serviceImpl struct {
	repository domain.MovieRepository
}

func CreateMovieService(r domain.MovieRepository) domain.MovieService {
	return &serviceImpl{r}
}

func (s serviceImpl) GetAll(ctx context.Context) ([]domain.Movie, error) {
	movie, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s serviceImpl) GetById(ctx context.Context, id int64) (*domain.Movie, error) {
	productId, err := s.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	return productId, nil
}

func (s serviceImpl) Create(ctx context.Context, name string, genre string, year int, award int) (*domain.Movie, error) {
	newProduct, err := s.repository.Create(ctx, name, genre, year, award)

	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s serviceImpl) UpdateAward(ctx context.Context, id int64, award int) (*domain.Movie, error) {
	productUpdate, err := s.repository.UpdateAward(ctx, id, award)

	if err != nil {
		return nil, err
	}

	return productUpdate, nil
}

func (s serviceImpl) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
