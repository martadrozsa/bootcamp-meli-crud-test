package user

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	MovieGenre string `json:"movie_genre"`
}

type UserRepository interface {
	GetAll() ([]*User, error)
	GetById(id int64) (*User, error)
	Create(name string, age int, movieGenre string) (*User, error)
	UpdateAge(id int64, age int) (*User, error)
	Delete(id int64)
}

type UserService interface {
	GetAll() ([]*User, error)
	GetById(id int64) (*User, error)
	Create(name string, age int, movieGenre string) (*User, error)
	UpdateAge(id int64, age int) (*User, error)
	Delete(id int64)
}
