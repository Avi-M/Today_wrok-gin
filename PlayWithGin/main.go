package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:name`
	Age  int    `json:age`
}

var Users []User

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", editUser)
		userRoutes.DELETE("/:id", deleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}

}
func GetUsers(c *gin.Context) {
	c.JSON(200, Users)
}
func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Invalid msg request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})
}
func editUser(c *gin.Context) {
	var id = c.Param("id")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid msg request body",
		})
		return
	}

	for i, user := range Users {
		if user.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})

}

func deleteUser(c *gin.Context) {
	var id = c.Param("id")
	for i, user := range Users {
		if user.ID == id {
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})
}
