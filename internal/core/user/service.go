package user

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUsers() ([]User, error) {
	var users []User

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUser(id string) (User, error) {
	var user User
	user, err := s.repo.FindByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) CreateUser(user User) (string, error) {
	isExist, err := s.repo.IsUserExist(user.Email)
	if err != nil {
		return "", err
	}

	if isExist {
		return "err user already exist", nil
	}

	if err := s.repo.CreateUser(user); err != nil {
		return "", err
	}

	return "success", nil
}
