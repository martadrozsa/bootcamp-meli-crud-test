package mysql

import (
	"context"
	"database/sql"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
)

type mysqlDBUserRepository struct {
	db *sql.DB
}

func (m *mysqlDBUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}

	rows, err := m.db.QueryContext(ctx, sqlGetAll)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User

		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.MovieGenre)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}

func (m *mysqlDBUserRepository) GetById(ctx context.Context, id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlDBUserRepository) Create(ctx context.Context, name string, age int, movieGenre string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlDBUserRepository) UpdateAge(ctx context.Context, id int64, age int) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlDBUserRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func CreateUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlDBUserRepository{db: db}
}
