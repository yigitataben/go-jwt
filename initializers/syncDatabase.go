package initializers

import "github.com/yigitataben/go-jwt/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
