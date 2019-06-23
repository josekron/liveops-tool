package user

//User represent a player with the total score
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	TotalScore int    `json:"totalScore"`
}
