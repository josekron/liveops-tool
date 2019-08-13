package user

import (
	"database/sql"
	"fmt"
	models "liveops-tool/user/models"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
)

type UserPostgrestDirectory struct {
}

func (u UserPostgrestDirectory) getConnection() *sql.DB {
	var password = os.Getenv("BD_PWD")
	var dbname = os.Getenv("BD_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	return db
}

func (u UserPostgrestDirectory) getUsers() []models.User {
	db := u.getConnection()
	defer db.Close()
	err := db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users = []models.User{}
	for rows.Next() {
		var id int
		var name string
		var country string
		var total_score int
		err = rows.Scan(&name, &id, &country, &total_score)
		checkErr(err)
		fmt.Println("id | name | country | total_score ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, country, total_score)
		user := models.User{
			ID:         id,
			Name:       strings.TrimSpace(name),
			Country:    strings.TrimSpace(country),
			TotalScore: total_score,
		}

		users = append(users, user)
	}
	return users
}

func (u UserPostgrestDirectory) SearchUsers(numUsers, minScore, maxScore int) []models.User {
	db := u.getConnection()
	defer db.Close()
	err := db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM users WHERE total_score >= $1 AND total_score <= $2 LIMIT $3", minScore, maxScore, numUsers)
	if err != nil {
		panic(err)
	}

	var users = []models.User{}
	for rows.Next() {
		var id int
		var name string
		var country string
		var total_score int
		err = rows.Scan(&name, &id, &country, &total_score)
		checkErr(err)
		fmt.Println("id | name | country | total_score ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, country, total_score)
		user := models.User{
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
