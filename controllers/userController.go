package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yigitataben/go-jwt/initializers"
	"github.com/yigitataben/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	var bodyStruct struct {
		UserName     string `json:"user_name"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		EmailAddress string `json:"email_address"`
		UserPassword string `json:"user_password"`
	}

	if err := c.BindJSON(&bodyStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body."})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(bodyStruct.UserPassword), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password."})
		return
	}
	user := models.User{UserName: bodyStruct.UserName, FirstName: bodyStruct.FirstName, LastName: bodyStruct.LastName, EmailAddress: bodyStruct.EmailAddress, UserPassword: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user."})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var bodyStruct struct {
		EmailAddress string `json:"email_address"`
		UserPassword string `json:"user_password"`
	}

	if err := c.BindJSON(&bodyStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body."})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email_address = ?", bodyStruct.EmailAddress)

	if user.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(bodyStruct.UserPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password."})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token."})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
