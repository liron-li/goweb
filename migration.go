package main

import (
	"goweb/models"
)

func Migration() {
	models.DB.AutoMigrate(&models.User{})
}
