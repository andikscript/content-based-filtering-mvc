package service

import (
	"contentbasedfilteringmvc/cmd/database"
	"contentbasedfilteringmvc/cmd/model"
)

func AuthUser(user model.User) (bool, string) {
	isExist, result := database.ReadUser(user)

	if isExist {
		if isValid := user.Username == result.Username && user.Password == result.Password; isValid {
			return true, user.Username
		}

		return false, ""
	}

	return false, ""
}
