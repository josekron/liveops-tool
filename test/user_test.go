package test

import (
	"fmt"
	userDirectory "liveops-tool/user/dao"
	userService "liveops-tool/user/services"
	"testing"
)

const (
	numUsers = 1
	minScore = 1500
	maxScore = 2000
)

func TestGenerateUserListByScore(t *testing.T) {

	var userDirectoryService = userService.NewService(userDirectory.UserDirectory{})
	var users = userDirectoryService.GenerateUserListByScore(numUsers, minScore, maxScore)
	fmt.Printf("%v", users)

	if len(users) != 1 {
		t.Errorf("Expected users list length = %d, but it was %d instead.", numUsers, len(users))
	}

	if users[0].TotalScore < minScore || users[0].TotalScore > maxScore {
		t.Errorf("Expected total score between %d and %d , but it was %d instead.", minScore, maxScore, users[0].TotalScore)
	}

	users = userDirectoryService.GenerateUserListByScore(numUsers+1, 0, maxScore)
	fmt.Printf("%v", users)

	if len(users) != 2 {
		t.Errorf("Expected users list length = %d, but it was %d instead.", numUsers, len(users))
	}
}
