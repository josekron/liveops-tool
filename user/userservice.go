package user

type UserService struct {
	dir Directory
}

func NewService(dir Directory) *UserService {
	var service = &UserService{
		dir: dir,
	}
	return service
}

// GenerateUserListByScore returns a list of users for a tournament
func (u UserService) GenerateUserListByScore(numUsers, minScore, maxScore int) ([]User, error) {
	return u.dir.SearchUsers(numUsers, minScore, maxScore)
}
