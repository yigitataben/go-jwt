package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitataben/go-jwt/controllers"
	"github.com/yigitataben/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}
func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.Run()
}
