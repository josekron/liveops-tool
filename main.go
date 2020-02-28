package main

import (
	"database/sql"
	"fmt"
	"liveops-tool/user"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	hostDB = "localhost"
	portDB = 5432
	userDB = "postgres"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Users
	r.GET("/competition/users", func(c *gin.Context) {
		numUsers, err := strconv.Atoi(c.DefaultQuery("numusers", "6"))
		minScore, err := strconv.Atoi(c.DefaultQuery("minscore", "0"))
		maxScore, err := strconv.Atoi(c.DefaultQuery("maxscore", "2000"))
		fmt.Printf("numUsers: %d , minScore: %d , maxScore: %d \n", numUsers, minScore, maxScore)

		if err == nil {

			db := getConnection()

			var userDirectoryService = user.NewService(user.UserPostgrestDirectory{Database: db})
			var res, err = userDirectoryService.GenerateUserListByScore(numUsers, minScore, maxScore)

			if err != nil {
				c.JSON(503, gin.H{"error": "db error"})
			} else {
				c.JSON(http.StatusOK, gin.H{"users": res})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request not valid"})
		}

	})

	return r
}

func getConnection() *sql.DB {
	var pwdDB = os.Getenv("BD_PWD")
	var nameDB = os.Getenv("BD_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", hostDB, portDB, userDB, pwdDB, nameDB)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
