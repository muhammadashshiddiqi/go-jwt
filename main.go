package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test jwt",
		})
	})
	r.Run()
}
