package domain

import "context"

type Movie struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
	Award int    `json:"award"`
}

type MovieRepository interface {
	GetAll(ctx context.Context) (*[]Movie, error)
	GetById(ctx context.Context, id int64) (*Movie, error)
	Create(ctx context.Context, movie *Movie) (*Movie, error)
	UpdateAward(ctx context.Context, movie *Movie) (*Movie, error)
	Delete(ctx context.Context, id int64) error
}

type MovieService interface {
	GetAll(ctx context.Context) (*[]Movie, error)
	GetById(ctx context.Context, id int64) (*Movie, error)
	Create(ctx context.Context, movie *Movie) (*Movie, error)
	UpdateAward(ctx context.Context, id int64, award int) (*Movie, error)
	Delete(ctx context.Context, id int64) error
}
