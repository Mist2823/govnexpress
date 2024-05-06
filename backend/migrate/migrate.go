package main

import (
	"vnexpress/initializers"
	"vnexpress/models"
)

func main() {
	// connect db
	initializers.LoadEnvVariables()
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		panic(err)
	}
}
