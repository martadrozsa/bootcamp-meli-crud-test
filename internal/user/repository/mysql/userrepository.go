package mysql

import (
	"fmt"
	"github.com/martadrozsa/bootcamp-meli-crud-test/internal/user/domain"
)

var listUsers []domain.User
var lastId int64 = 1

type repositoryImpl struct {
}

func CreateUserRepository() domain.UserRepository {
	listUsers = []domain.User{}
	listUsers = append(listUsers)

	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll() ([]domain.User, error) {
	return listUsers, nil
}

func (r *repositoryImpl) GetById(id int64) (*domain.User, error) {
	for _, user := range listUsers {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("the user with the id %d was not found", id)
}

func (r *repositoryImpl) Create(name string, age int, movieGenre string) (*domain.User, error) {
	nextId := lastId
	newUser := domain.User{
		Id:         nextId,
		Name:       name,
		Age:        age,
		MovieGenre: movieGenre,
	}
	for _, user := range listUsers {
		if user.Name == name {
			return nil, fmt.Errorf("the user with name %s has already been registered", name)
		}
	}
	listUsers = append(listUsers, newUser)
	lastId += 1

	return &newUser, nil
}

func (r *repositoryImpl) UpdateAge(id int64, age int) (*domain.User, error) {

	var user domain.User
	update := false
	for i := range listUsers {
		if listUsers[i].Id == id {
			listUsers[i].Age = age
			update = true
			user = listUsers[i]
			break
		}
	}

	if !update {
		return nil, fmt.Errorf("the product with id %d was not found", id)
	}
	return &user, nil

}

func (r *repositoryImpl) Delete(id int64) error {
	deleted := false
	var index int

	for i := range listUsers {
		if listUsers[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("the product with id %d was not found", id)
	}

	listUsers = append(listUsers[:index], listUsers[index+1:]...)

	return nil
}
