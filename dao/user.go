// Package dao 数据库操作
package dao

import (
	"go.uber.org/zap"
	"go_chat/global"
	"go_chat/models"
)

// GetUserList 获取用户列表
func GetUserList() (list []*models.UserBasic) {
	global.DB.Find(&list)
	return
}

// FindUserByName GetUserByName 通过用户名查询（登录时使用）
func FindUserByName(name string) (user *models.UserBasic, code int) {
	user = &models.UserBasic{}
	if tx := global.DB.Where("name= ?", name).First(user); tx.RowsAffected == 0 {
		code = 1002
	}
	return
}

// UserExist 判断用户名是否已被占用（注册时使用）
func UserExist(name string) (code int) {
	user := &models.UserBasic{}
	if tx := global.DB.Where("name= ?", name).First(user); tx.RowsAffected == 1 {
		code = 1001
	}
	return
}

// FindUserById 通过Id查询用户
func FindUserById(id uint) (user *models.UserBasic, code int) {
	user = &models.UserBasic{}
	if tx := global.DB.Where(id).First(user); tx.RowsAffected == 0 {
		code = 1002
	}
	return
}

// FindUserByPhone 通过电话号码查询用户
func FindUserByPhone(phone string) (user *models.UserBasic, code int) {
	user = &models.UserBasic{}
	if tx := global.DB.Where("phone= ?", phone).First(user); tx.RowsAffected == 0 {
		code = 1002
	}
	return
}

// FindUserByEmail 通过邮箱查询用户
func FindUserByEmail(email string) (user *models.UserBasic, code int) {
	user = &models.UserBasic{}
	if tx := global.DB.Where("email= ?", email).First(user); tx.RowsAffected == 0 {
		code = 1002
	}
	return
}

// CreateUser 创建用户
func CreateUser(user *models.UserBasic) (code int) {
	if tx := global.DB.Create(user); tx.RowsAffected == 0 {
		zap.S().Info("failed to create user")
		code = 500
	}
	return
}

// UpdateUser 更新用户
func UpdateUser(user *models.UserBasic) (code int) {
	tx := global.DB.Model(user).Updates(user)
	if tx.RowsAffected == 0 {
		zap.S().Info("failed to update user")
		code = 500
	}
	return
}

// DeleteUser 删除用户
func DeleteUser(user *models.UserBasic) (code int) {
	if tx := global.DB.Delete(user); tx.RowsAffected == 0 {
		zap.S().Info("failed to delete user")
		code = 500
	}
	return
}

// FindUserByNameAndPwd 通过用户名和密码查找用户（用户登录）
func FindUserByNameAndPwd(name, password string) (user *models.UserBasic, code int) {
	user = &models.UserBasic{}
	tx := global.DB.Where("name= ? and pass_word= ?", name, password).First(user)
	if tx.RowsAffected == 0 {
		code = 1003
	}
	return
}
