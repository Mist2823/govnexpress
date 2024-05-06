package initializers

import "vnexpress/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
