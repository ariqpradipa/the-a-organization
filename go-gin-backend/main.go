package main

import (
	"bookweb/controllers"
	middleware "bookweb/middlewares"
	"bookweb/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the bookweb API",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/books", controllers.FindBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PUT("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/user", middleware.JWTAuthMiddleware(), controllers.CurrentUser)
	}
	// health check
	r.GET("/health/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "true",
		})
	})

	r.Run()
}
