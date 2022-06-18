package user

type serviceImpl struct {
}

func CreateUserService() UserService {
	return &serviceImpl{}
}

func (s *serviceImpl) GetAll() ([]*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) GetById(id int64) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Create(name string, age int, movieGenre string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) UpdateAge(id int64, age int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Delete(id int64) {
	//TODO implement me
	panic("implement me")
}
