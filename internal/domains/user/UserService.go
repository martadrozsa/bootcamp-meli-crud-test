package user

type serviceImpl struct {
	repository UserRepository
}

func CreateUserService(r UserRepository) UserService {
	return &serviceImpl{r}
}

func (s *serviceImpl) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *serviceImpl) GetById(id int64) (*User, error) {
	user, err := s.repository.GetById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *serviceImpl) Create(name string, age int, movieGenre string) (*User, error) {
	newUser, err := s.repository.Create(name, age, movieGenre)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *serviceImpl) UpdateAge(id int64, age int) (*User, error) {
	return s.repository.UpdateAge(id, age)
}

func (s *serviceImpl) Delete(id int64) error {
	return s.repository.Delete(id)
}
