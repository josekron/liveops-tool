package test

import (
	"fmt"
	"liveops-tool/user"
	"testing"
)

const (
	numUsers = 1
	minScore = 1500
	maxScore = 2000
)

func TestGenerateUserListByScore(t *testing.T) {

	var userDirectoryService = user.NewService(user.UserDirectory{})
	var users, err = userDirectoryService.GenerateUserListByScore(numUsers, minScore, maxScore)
	fmt.Printf("%v", users)

	if len(users) != 1 || err != nil {
		t.Errorf("Expected users list length = %d, but it was %d instead.", numUsers, len(users))
	}

	if users[0].TotalScore < minScore || users[0].TotalScore > maxScore {
		t.Errorf("Expected total score between %d and %d , but it was %d instead.", minScore, maxScore, users[0].TotalScore)
	}

	users, err = userDirectoryService.GenerateUserListByScore(numUsers+1, 0, maxScore)
	fmt.Printf("%v", users)

	if len(users) != 2 || err != nil {
		t.Errorf("Expected users list length = %d, but it was %d instead.", numUsers, len(users))
	}
}
