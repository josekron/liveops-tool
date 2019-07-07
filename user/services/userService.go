package user

import (
	dao "liveops-tool/user/dao"
	models "liveops-tool/user/models"
)

type UserService struct {
	dir dao.Directory
}

func NewService(dir dao.Directory) *UserService {
	var service = &UserService{
		dir: dir,
	}
	return service
}

// GenerateUserListByScore returns a list of users for a tournament
func (u UserService) GenerateUserListByScore(numUsers, minScore, maxScore int) []models.User {
	return u.dir.SearchUsers(numUsers, minScore, maxScore)
}
