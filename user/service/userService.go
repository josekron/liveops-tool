package user

import (
	"fmt"
	userModel "liveops-tool/user/model"
)

type Directory interface {
	GetUsers() []userModel.User
	SearchUsers(numUsers, minScore, maxScore int) []userModel.User
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

func (u UserService) GetUsers() []userModel.User {
	return u.dir.GetUsers()
}

func (u UserService) SearchUsers(numUsers, minScore, maxScore int) []userModel.User {
	return u.dir.SearchUsers(numUsers, minScore, maxScore)
}

// // GenerateUserListByScore returns a list of users for a tournament
// func GenerateUserListByScore(numUsers, minScore, maxScore int) []userModel.User {
// 	return searchUsers(numUsers, minScore, maxScore)
// }
