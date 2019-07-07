package user

import (
	models "liveops-tool/user/models"
)

type Directory interface {
	SearchUsers(numUsers, minScore, maxScore int) []models.User
}
