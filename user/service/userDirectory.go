package user

import (
	"fmt"
	userModel "liveops-tool/user/model"
)

var users = []userModel.User{}

func init() {
	initUsers()
}

func searchUsers(numUsers, minScore, maxScore int) []userModel.User {
	return filterUsersByScore(numUsers, minScore, maxScore)
}

func filterUsersByScore(numUsers, minScore, maxScore int) []userModel.User {
	var filteredUsers = []userModel.User{}
	var countUsers = 0
	if len(users) < numUsers {
		numUsers = len(users)
	}
	for i := range users {
		if users[i].TotalScore >= minScore && users[i].TotalScore <= maxScore {
			filteredUsers = append(filteredUsers, users[i])
			countUsers++
		}

		if countUsers >= numUsers {
			break
		}
	}
	return filteredUsers
}

func initUsers() {
	user1 := userModel.User{ID: "1", Name: "John", Country: "UK", TotalScore: 500}
	user2 := userModel.User{ID: "2", Name: "Frank", Country: "ES", TotalScore: 1500}
	user3 := userModel.User{ID: "3", Name: "Bob", Country: "UK", TotalScore: 2000}
	user4 := userModel.User{ID: "4", Name: "Anna", Country: "FR", TotalScore: 3000}
	user5 := userModel.User{ID: "5", Name: "Marta", Country: "UK", TotalScore: 1400}
	user6 := userModel.User{ID: "6", Name: "Katy", Country: "UK", TotalScore: 1800}
	user7 := userModel.User{ID: "7", Name: "Pepe", Country: "BR", TotalScore: 400}
	user8 := userModel.User{ID: "8", Name: "Maria", Country: "ES", TotalScore: 1000}
	user9 := userModel.User{ID: "9", Name: "Emma", Country: "UK", TotalScore: 1200}
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)
	users = append(users, user4)
	users = append(users, user5)
	users = append(users, user6)
	users = append(users, user7)
	users = append(users, user8)
	users = append(users, user9)
	fmt.Printf("users: %+v\n", users)
}
