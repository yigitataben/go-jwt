package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitataben/go-jwt/controllers"
	"github.com/yigitataben/go-jwt/initializers"
	"github.com/yigitataben/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)

	r.Run()
}
