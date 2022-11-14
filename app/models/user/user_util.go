package user

import (
	"go-web/pkg/database"
	"os/user"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(user.User{}).Where("email=?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 Phone 已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(user.User{}).Where("phone=?", phone).Count(&count)
	return count > 0
}
