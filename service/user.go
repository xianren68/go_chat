// Package service 对外api
package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_chat/common"
	"go_chat/dao"
	"go_chat/middleware"
	"go_chat/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// List 获取用户列表
func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		//ctx.JSON(200, gin.H{"code": -1, "msg": "获取用户列表失败"})
		errReply(ctx, "failed to get user list")
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// Login 登录
func Login(ctx *gin.Context) {
	info := make(map[string]string)
	ctx.BindJSON(&info)
	name, pwd := info["name"], info["password"]
	if name == "" || pwd == "" {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "username or password is null",
		//})
		errReply(ctx, "username or password is null")
		return
	}
	user, err := dao.FindUserByName(name)
	if err != nil {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	// 加密密码
	pwd = common.SaltPassword(pwd, user.Salt)
	// 查询是否存在
	user, err = dao.FindUserByNameAndPwd(name, pwd)
	if err != nil {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	// 添加token
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		zap.S().Info(err)
	}
	ctx.Header("Authorization", "Bearer "+token)
	okReply(ctx)
}

// NewUser 添加用户
func NewUser(ctx *gin.Context) {
	info := make(map[string]string)
	err := ctx.BindJSON(&info)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "failed to read arguments",
		//})
		errReply(ctx, "failed to read arguments")
		return
	}
	name := info["name"]
	password := info["password"]
	// 确认密码
	rePassword := info["identity"]
	if name == "" || password == "" || rePassword == "" {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "username and password can not be empty",
		//})
		errReply(ctx, "username and password can not be empty")
		return
	}
	// 判断两次密码是否一致
	if password != rePassword {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "two passwords are inconsistent",
		//})
		errReply(ctx, "two passwords are inconsistent")
		return
	}
	// 判断用户名是否被占用
	err = dao.UserExist(name)
	if err != nil {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	// 生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())
	// 对密码加密
	pwd := common.SaltPassword(password, salt)
	t := time.Now()
	err = dao.CreateUser(&models.UserBasic{
		Name:          name,
		PassWord:      pwd,
		Salt:          salt,
		LoginTime:     &t,
		HeartBeatTime: &t,
		LoginOutTime:  &t,
	})
	if err != nil {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	okReply(ctx)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	err := ctx.BindJSON(user)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "failed to read arguments",
		//})
		errReply(ctx, "failed to read arguments")
		return
	}
	// 判断要更新的用户名是否已经存在
	err = dao.UserExist(user.Name)
	if err != nil {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	// 如果存在密码的改变，继续加密
	if user.PassWord != "" {
		salt := fmt.Sprintf("%d", rand.Int31())
		pwd := common.SaltPassword(user.PassWord, salt)
		user.PassWord = pwd
		user.Salt = salt
	}
	// 更新
	err = dao.UpdateUser(user)
	if err != nil {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	okReply(ctx)
}

// DeleteUser 注销账号
func DeleteUser(ctx *gin.Context) {
	Id := ctx.PostForm("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		zap.S().Info("id is not compliant")
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  "id is not compliant",
		//})
		errReply(ctx, "id is not compliant")
		return
	}
	user := &models.UserBasic{}
	user.ID = uint(id)
	err = dao.DeleteUser(user)
	if err != nil {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(ctx, err.Error())
		return
	}
	okReply(ctx)
}

// 成功回复
func okReply(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

// 失败回复
func errReply(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  err,
	})
}
