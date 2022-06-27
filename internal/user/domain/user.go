package domain

import "context"

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	MovieGenre string `json:"movie_genre"`
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, name string, age int, movieGenre string) (*User, error)
	UpdateAge(ctx context.Context, id int64, age int) (*User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, name string, age int, movieGenre string) (*User, error)
	UpdateAge(ctx context.Context, id int64, age int) (*User, error)
	Delete(ctx context.Context, id int64) error
}
