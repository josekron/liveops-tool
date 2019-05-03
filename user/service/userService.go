package user

import (
	userModel "liveops-tool/user/model"
)

// GenerateUserListByScore returns a list of users for a tournament
func GenerateUserListByScore(numUsers, minScore, maxScore int) []userModel.User {
	return searchUsers(numUsers, minScore, maxScore)
}
