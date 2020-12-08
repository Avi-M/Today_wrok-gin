package main

import (
	"ginCamp/controller"
	"ginCamp/service"
	"io"
	"os"

	"ginCamp/middlewares"

	"github.com/gin-gonic/gin"

	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
func main() {

	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/getvideos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/postvideos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(400, gin.H{"msg": err.Error()})
			} else {
				c.JSON(200, gin.H{
					"message": "video input is valid",
				})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
