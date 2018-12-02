package models

type User struct {
	Model
	Username string
	Password string
}

func CheckAuth(username, password string) bool {
	var user User
	DB.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
