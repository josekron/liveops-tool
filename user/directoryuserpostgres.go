package user

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type UserPostgrestDirectory struct {
	Database *sql.DB
}

func (u UserPostgrestDirectory) getUsers() []User {

	rows, err := u.Database.Query("SELECT * FROM users")
	checkErr(err)

	var users = []User{}
	for rows.Next() {
		var id int
		var name string
		var country string
		var total_score int
		err = rows.Scan(&name, &id, &country, &total_score)
		checkErr(err)
		fmt.Println("id | name | country | total_score ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, country, total_score)
		user := User{
			ID:         id,
			Name:       strings.TrimSpace(name),
			Country:    strings.TrimSpace(country),
			TotalScore: total_score,
		}

		users = append(users, user)
	}
	return users
}

func (u UserPostgrestDirectory) SearchUsers(numUsers, minScore, maxScore int) []User {

	rows, err := u.Database.Query("SELECT * FROM users WHERE total_score >= $1 AND total_score <= $2 LIMIT $3", minScore, maxScore, numUsers)
	checkErr(err)

	var users = []User{}
	for rows.Next() {
		var id int
		var name string
		var country string
		var total_score int
		err = rows.Scan(&name, &id, &country, &total_score)
		checkErr(err)
		fmt.Println("id | name | country | total_score ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, country, total_score)
		user := User{
			ID:         id,
			Name:       strings.TrimSpace(name),
			Country:    strings.TrimSpace(country),
			TotalScore: total_score,
		}

		users = append(users, user)
	}
	return users
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
