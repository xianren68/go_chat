package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/common"
	"go_chat/dao"
	"net/http"
)

// FriendList 获取好友列表
func FriendList(c *gin.Context) {
	// 取出jwt中的id
	value, _ := c.Get("userId")
	// 获取数据
	list := dao.FriendList(value.(uint))
	//if err != nil {
	//	//c.JSON(http.StatusOK, gin.H{
	//	//	"code": -1,
	//	//	"msg":  err.Error(),
	//	//})
	//	errReply(c, err.Error())
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": list,
	})
}

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	// 取出jwt中的id
	value, _ := c.Get("userId")
	x := make(map[string]string)
	err := c.BindJSON(&x)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	name := x["name"]
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "用户名不能为空",
		})
		//errReply(c, "userName is null")
		return
	}
	code := dao.AddFriendByName(value.(uint), name)
	if code != 0 {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(c, code)
		return
	}
	common.OkReply(c)

}
