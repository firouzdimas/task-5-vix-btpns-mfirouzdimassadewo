package main

import (
	"task-5-vix-fullstack/database"
	"task-5-vix-fullstack/models"
	"task-5-vix-fullstack/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
