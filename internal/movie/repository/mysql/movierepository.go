package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/movie/domain"
)

type mysqlDBMovieRepository struct {
	db *sql.DB
}

func CreateMovieRepository(db *sql.DB) domain.MovieRepository {
	return &mysqlDBMovieRepository{db: db}
}

func (m *mysqlDBMovieRepository) GetAll(ctx context.Context) ([]domain.Movie, error) {
	movies := []domain.Movie{}

	rows, err := m.db.QueryContext(ctx, sqlGetAll)

	if err != nil {
		return movies, err
	}

	defer rows.Close()

	for rows.Next() {
		var movie domain.Movie

		err := rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.Year, &movie.Award)
		if err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m *mysqlDBMovieRepository) GetById(ctx context.Context, id int64) (*domain.Movie, error) {
	var movie domain.Movie

	rows, err := m.db.QueryContext(ctx, sqlGetById, id)

	if err != nil {
		return nil, err
	}

	if rows.Next() != true {
		return nil, errors.New("not found")
	}

	err = rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.Year, &movie.Award)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &movie, nil

}

func (m *mysqlDBMovieRepository) Create(ctx context.Context, name string, genre string, year int, award int) (*domain.Movie, error) {
	movie, err := m.db.ExecContext(
		ctx,
		sqlCreate,
		name,
		genre,
		year,
		award,
	)

	if err != nil {
		return nil, err
	}

	newMovieId, err := movie.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &domain.Movie{
		Id:    newMovieId,
		Name:  name,
		Genre: genre,
		Year:  year,
		Award: award,
	}, nil
}

func (m *mysqlDBMovieRepository) UpdateAward(ctx context.Context, id int64, award int) (*domain.Movie, error) {
	_, err := m.db.ExecContext(
		ctx,
		sqlUpdateAward,
		award,
		id,
	)

	if err != nil {
		return nil, err
	}

	movieUpdate, err := m.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return movieUpdate, nil
}

func (m *mysqlDBMovieRepository) Delete(ctx context.Context, id int64) error {
	_, err := m.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	return nil
}
