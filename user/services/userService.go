package user

import (
	"fmt"
	models "liveops-tool/user/models"
)

type Directory interface {
	GetUsers() []models.User
	SearchUsers(numUsers, minScore, maxScore int) []models.User
}

type UserService struct {
	dir Directory
}

func NewService(dir Directory) *UserService {
	var service = &UserService{
		dir: dir,
	}
	fmt.Println(service.dir.GetUsers())
	return service
}

func (u UserService) GetUsers() []models.User {
	return u.dir.GetUsers()
}

// GenerateUserListByScore returns a list of users for a tournament
func (u UserService) GenerateUserListByScore(numUsers, minScore, maxScore int) []models.User {
	return u.dir.SearchUsers(numUsers, minScore, maxScore)
}
