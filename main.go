package main

import (
	"fmt"
	userService "liveops-tool/user/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

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
			var userDirectoryService = userService.NewService(userService.UserDirectory{})
			var res = userDirectoryService.SearchUsers(numUsers, minScore, maxScore)
			c.JSON(http.StatusOK, gin.H{"users": res})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"users": "no valid"})
		}

	})

	// Test Enpoints

	// Ping test
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "pong")
	// })

	// Get user value
	// r.GET("/user", func(c *gin.Context) {
	// 	res := []string{}
	// 	for k, v := range db {
	// 		res = append(res, k, v)
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"users": res})
	// })

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	fmt.Printf(user + "\n")
	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
