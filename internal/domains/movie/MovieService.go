package movie

type serviceImpl struct {
	repository MovieRepository
}

func CreateMovieService(r MovieRepository) MovieService {
	return &serviceImpl{r}
}

func (s *serviceImpl) GetAll() ([]*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) GetById(id int64) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Create(name string, genre string, year int, award string) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) UpdateAward(id int64, award string) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceImpl) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
