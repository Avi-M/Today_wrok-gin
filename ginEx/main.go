package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/test/:name", func(c *gin.Context) {
		var val = c.Param("name")
		c.JSON(200, gin.H{
			"message": val,
		})
	})

	route.POST("/add", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{
			"person": person,
		})
	})

	route.Run(":8085")
}
