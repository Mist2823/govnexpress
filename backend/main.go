package main

import (
	"vnexpress/controllers"
	"vnexpress/initializers"
	"vnexpress/middleware"
	"vnexpress/scraper"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	corsMiddleware := cors.Default()

	r.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Require, controllers.Validate)

	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("failed to connect database")
	}

	scraper.ScrapeAndStore(db)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/post/:id", controllers.PostsDelete)

	r.Run()
}
