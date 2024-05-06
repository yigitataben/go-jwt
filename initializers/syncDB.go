package initializers

import "github.com/yigitataben/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})

}
