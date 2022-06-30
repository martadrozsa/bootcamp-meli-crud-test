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
	return &mysqlDBMovieRepository{
		db: db}
}

func (m *mysqlDBMovieRepository) GetAll(ctx context.Context) (*[]domain.Movie, error) {

	var movies []domain.Movie

	rows, err := m.db.QueryContext(ctx, sqlGetAll)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var movie domain.Movie

		err := rows.Scan(
			&movie.Id,
			&movie.Name,
			&movie.Genre,
			&movie.Year,
			&movie.Award)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}
	return &movies, nil
}

func (m *mysqlDBMovieRepository) GetById(ctx context.Context, id int64) (*domain.Movie, error) {

	rows := m.db.QueryRowContext(ctx, sqlGetById, id)

	var movie domain.Movie

	err := rows.Scan(
		&movie.Id,
		&movie.Name,
		&movie.Genre,
		&movie.Year,
		&movie.Award)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrIDNotFound
	}

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *mysqlDBMovieRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {

	movieResult, err := m.db.ExecContext(
		ctx,
		sqlCreate,
		&movie.Name,
		&movie.Genre,
		&movie.Year,
		&movie.Award,
	)

	if err != nil {
		return nil, err
	}

	lastId, err := movieResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	movie.Id = lastId

	return movie, nil
}

func (m *mysqlDBMovieRepository) UpdateAward(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {

	movieResult, err := m.db.ExecContext(
		ctx,
		sqlUpdateAward,
		&movie.Award,
		&movie.Id,
	)

	if err != nil {
		return nil, err
	}

	affectedRows, err := movieResult.RowsAffected()
	if affectedRows == 0 {
		return nil, domain.ErrIDNotFound
	}

	if err != nil {
		return nil, err
	}

	movieUpdate, err := m.GetById(ctx, movie.Id)
	if err != nil {
		return nil, err
	}

	return movieUpdate, nil
}

func (m *mysqlDBMovieRepository) Delete(ctx context.Context, id int64) error {

	result, err := m.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrIDNotFound
	}

	return nil
}
