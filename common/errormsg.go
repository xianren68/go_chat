package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一些常量，用做错误状态码
const (
	SUCCESS = 0
	ERROR   = 500
	// ErrorUserUsed  关于用户数据的错误 1000
	ErrorUserUsed        = 1001
	ErrorUserNotExist    = 1002
	ErrorPasswordWrong   = 1003
	ErrorUserNoRight     = 1004
	ErrorTokenNotExist   = 1005
	ErrorTokenWrong      = 1006
	ErrorTokenRuntime    = 1007
	ErrorTokenTypeWrong  = 1008
	IsSameUser           = 1009
	RepeatAddFriend      = 1010
	ErrorGroupExist      = 2001
	ErrorRepeatJoinGroup = 2003
	ErrorGroupNotExist   = 2004
)

// 建立错误信息与错误码的映射表
var codeMsg = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	ErrorUserUsed:        "用户名被占用",
	ErrorUserNotExist:    "用户不存在",
	ErrorUserNoRight:     "该用户无权限",
	ErrorPasswordWrong:   "密码有误",
	ErrorTokenNotExist:   "token不存在,请重新登录",
	ErrorTokenWrong:      "token有误,请重新登录",
	ErrorTokenRuntime:    "token过期,请重新登录",
	ErrorTokenTypeWrong:  "token格式有误,请重新登录",
	IsSameUser:           "不能添加自己为好友",
	RepeatAddFriend:      "已存在此好友",
	ErrorGroupExist:      "群组已存在",
	ErrorRepeatJoinGroup: "不能重复加群",
	ErrorGroupNotExist:   "不存在此群组",
}

// GetMsg 给外面的接口,通过错误码获取msg
func GetMsg(code int) string {
	return codeMsg[code]
}

// OkReply 成功回复
func OkReply(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

// ErrReply 失败回复
func ErrReply(
	ctx *gin.Context,
	code int,
) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  GetMsg(code),
	})
}
