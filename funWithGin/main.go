package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func postHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}
func queryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func pathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func main() {
	r := gin.Default()
	r.GET("/ping", sayHello)
	r.POST("/", postHomePage)
	r.GET("/query", queryStrings) // run http://localhost:8080/query/?name=Avi&age=23
	r.GET("/path/:name/:age", pathParameters)
	r.Run()
}
