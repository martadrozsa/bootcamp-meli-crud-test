package movie

type repositoryImpl struct {
}

func CreateMovieRepository() MovieRepository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll() ([]*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) GetById(id int64) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Create(name string, genre string, year int, award string) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) UpdateAward(id int64, award string) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryImpl) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
