package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/dao"
	"net/http"
)

// FriendList 获取好友列表
func FriendList(c *gin.Context) {
	// 取出jwt中的id
	value, _ := c.Get("userId")
	// 获取数据
	list, err := dao.FriendList(value.(uint))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
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
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "userName is null",
		})
		return
	}
	code, err := dao.AddFriendByName(value.(uint), name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}
	okReply(c)

}
