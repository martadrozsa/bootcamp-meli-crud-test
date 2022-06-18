package movie

type Movie struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
	Award string `json:"award"`
}

type MovieRepository interface {
	GetAll() ([]*Movie, error)
	GetById(id int64) (*Movie, error)
	Create(name string, genre string, year int, award string) (*Movie, error)
	UpdateAward(id int64, award string) (*Movie, error)
	Delete(id int64) error
}

type MovieService interface {
	GetAll() ([]*Movie, error)
	GetById(id int64) (*Movie, error)
	Create(name string, genre string, year int, award string) (*Movie, error)
	UpdateAward(id int64, award string) (*Movie, error)
	Delete(id int64) error
}
