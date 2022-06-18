package user

import "fmt"

var listUsers []User
var lastId int64 = 1

type repositoryImpl struct {
}

func CreateUserRepository() UserRepository {
	listUsers = []User{}
	listUsers = append(listUsers)

	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll() ([]User, error) {
	return listUsers, nil
}

func (r *repositoryImpl) GetById(id int64) (*User, error) {
	for _, user := range listUsers {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("the user with the id %d was not found", id)
}

func (r *repositoryImpl) Create(name string, age int, movieGenre string) (*User, error) {
	nextId := lastId
	newUser := User{
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

func (r *repositoryImpl) UpdateAge(id int64, age int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Delete(id int64) {
	//TODO implement me
	panic("implement me")
}
