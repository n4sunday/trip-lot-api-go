package user

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id string) error
	IsUserExist(email string) (bool, error)
}

type UserService interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user User) (string, error)
}
