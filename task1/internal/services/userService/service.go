package userService

type UserService struct {
	repo *userRepoistory
}

func NewUserService(repo *userRepoistory) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PostUser(user User) error {
	return s.repo.PostUser(user)
}

func (s *UserService) PatchUserByID(userID int, user User) error {
	return s.repo.PatchUserByID(userID, user)
}

func (s *UserService) DeleteUserByID(userID int) error {
	return s.repo.DeleteUserByID(userID)
}
