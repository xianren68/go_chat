// Package service 对外api
package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_chat/common"
	"go_chat/dao"
	"go_chat/global"
	"go_chat/middleware"
	"go_chat/models"
	"go_chat/upload"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// List 获取用户列表
func List(ctx *gin.Context) {
	list := dao.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": list,
	})
}

// Login 登录
func Login(ctx *gin.Context) {
	info := make(map[string]string)
	_ = ctx.BindJSON(&info)
	name, pwd := info["name"], info["password"]
	if name == "" || pwd == "" {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "用户名或密码为空",
		})
		return
	}
	user, code := dao.FindUserByName(name)
	if code != 0 {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	// 加密密码
	pwd = common.SaltPassword(pwd, user.Salt)
	// 查询是否存在
	user, code = dao.FindUserByNameAndPwd(name, pwd)
	if code != 0 {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	// 添加token
	token, code := middleware.GenerateToken(user.ID)
	if code != 0 {
		return
	}
	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{
		"code":     0,
		"userInfo": user,
	})
}

// NewUser 添加用户
func NewUser(ctx *gin.Context) {
	info := make(map[string]string)
	err := ctx.BindJSON(&info)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数错误",
		})
		return
	}
	name := info["name"]
	password := info["password"]
	// 确认密码
	rePassword := info["identity"]
	if name == "" || password == "" || rePassword == "" {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "用户名和密码不能为空",
		})
		return
	}
	// 判断两次密码是否一致
	if password != rePassword {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "两次密码不一致",
		})
		return
	}
	// 判断用户名是否被占用
	code := dao.UserExist(name)
	if code != 0 {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	// 生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())
	// 对密码加密
	pwd := common.SaltPassword(password, salt)
	t := time.Now()
	code = dao.CreateUser(&models.UserBasic{
		Name:          name,
		PassWord:      pwd,
		Salt:          salt,
		LoginTime:     &t,
		HeartBeatTime: &t,
		LoginOutTime:  &t,
	})
	if code != 0 {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	common.OkReply(ctx)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	err := ctx.BindJSON(user)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数错误",
		})
		return
	}
	// 判断要更新的用户名是否已经存在
	code := dao.UserExist(user.Name)
	if code != 0 {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
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
	code = dao.UpdateUser(user)
	if code != 0 {
		//ctx.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	common.OkReply(ctx)
}

// UpdateImg 更新头像
func UpdateImg(ctx *gin.Context) {
	id, _ := ctx.Get("userId")
	Id := id.(uint)
	file, fileHeader, _ := ctx.Request.FormFile("file")
	// 获取文件长度
	fileSize := fileHeader.Size
	url, code := upload.UploadImg(file, fileSize)
	if code != 0 {
		common.ErrReply(ctx, code)
		return
	}
	user := &models.UserBasic{}
	user.ID = Id
	user.Avatar = url
	// 更新用户信息
	code = dao.UpdateUser(user)
	if code != 0 {
		common.ErrReply(ctx, code)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "ok",
		"data": url,
	})
}

// DeleteUser 注销账号
func DeleteUser(ctx *gin.Context) {
	Id := ctx.PostForm("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		zap.S().Info("id is not compliant")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请求参数有误",
		})
		return
	}
	user := &models.UserBasic{}
	user.ID = uint(id)
	code := dao.DeleteUser(user)
	if code != 0 {
		//ctx.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(ctx, code)
		return
	}
	common.OkReply(ctx)
}

// SendUserMsg 发送消息
func SendUserMsg(c *gin.Context) {
	dao.Chat(c)
}

// GetUnread 获取用户未读消息
func GetUnread(c *gin.Context) {
	value, _ := c.Get("userId")
	id := value.(uint)
	data := getUnreadList(id)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

// 获取redis未读消息列表
func getUnreadList(id uint) []string {
	ctx := context.Background()
	key := "list" + strconv.Itoa(int(id))
	count, _ := global.RedisDB.LLen(ctx, key).Result()
	result, err := global.RedisDB.LPopCount(ctx, key, int(count)).Result()
	if err != nil {
		zap.S().Info(err)
	}
	return result
}

// FindUserByName 通过用户名查找用户
func FindUserByName(c *gin.Context) {
	value := c.Query("name")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "failed to get name",
		})
		return
	}
	user, code := dao.FindUserByName(value)
	if code != 0 {
		common.ErrReply(c, code)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": user,
	})

}
