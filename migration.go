package main

import (
	"goweb/models"
	"qor-example/config/db"
)

func Migration() {
	db.DB.AutoMigrate(&models.User{})
}
