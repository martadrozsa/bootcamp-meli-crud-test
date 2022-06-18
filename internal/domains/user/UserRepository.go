package user

type repositoryImpl struct {
}

func CreateUserRepository() UserRepository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll() ([]*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) GetById(id int64) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Create(name string, age int, movieGenre string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) UpdateAge(id int64, age int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Delete(id int64) {
	//TODO implement me
	panic("implement me")
}
