package user

import (
	models "liveops-tool/user/models"
)

type Directory interface {
	SearchUsers(numUsers, minScore, maxScore int) []models.User
}

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
func (u UserService) GenerateUserListByScore(numUsers, minScore, maxScore int) []models.User {
	return u.dir.SearchUsers(numUsers, minScore, maxScore)
}
