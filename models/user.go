package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(User{UserName: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
