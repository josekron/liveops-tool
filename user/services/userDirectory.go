package user

import (
	models "liveops-tool/user/models"
)

type UserDirectory struct {
}

func (u UserDirectory) GetUsers() []models.User {
	var users = []models.User{
		{
			ID:         "1",
			Name:       "John",
			Country:    "UK",
			TotalScore: 500,
		},
		{
			ID:         "2",
			Name:       "Frank",
			Country:    "ES",
			TotalScore: 1500,
		},
		{
			ID:         "3",
			Name:       "Bob",
			Country:    "UK",
			TotalScore: 2000,
		},
		{
			ID:         "4",
			Name:       "Anna",
			Country:    "FR",
			TotalScore: 3000,
		},
	}

	return users
}

func (u UserDirectory) SearchUsers(numUsers, minScore, maxScore int) []models.User {
	return u.filterUsersByScore(numUsers, minScore, maxScore)
}

func (u UserDirectory) filterUsersByScore(numUsers, minScore, maxScore int) []models.User {
	var users = u.GetUsers()
	var filteredUsers = []models.User{}
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
