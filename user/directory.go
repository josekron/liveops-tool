package user

type Directory interface {
	SearchUsers(numUsers, minScore, maxScore int) ([]User, error)
}
