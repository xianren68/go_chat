package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_chat/dao"
	"go_chat/models"
	"net/http"
)

// NewGroup 新建群聊
func NewGroup(c *gin.Context) {
	value, _ := c.Get("userId")
	group := &models.Community{}
	err := c.BindJSON(group)
	// 群名为空
	if group.Name == "" {
		//c.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "group name cannot be empty",
		//})
		errReply(c, "group name cannot be empty")
		return
	}

	group.OwnerId = value.(uint)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		//c.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "failed to read arguments",
		//})
		errReply(c, "failed to read arguments")
		return
	}
	err = dao.CreateCommunity(group)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(c, err.Error())
		return
	}
	okReply(c)
}

// GetGroupList 获取群列表
func GetGroupList(c *gin.Context) {
	value, _ := c.Get("userId")
	group, err := dao.GetCommunities(value.(uint))
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": group,
	})
}

// JoinGroup 加入群聊
func JoinGroup(c *gin.Context) {
	value, _ := c.Get("userId")
	name := c.PostForm("groupName")
	if name == "" {
		//c.JSON(200, gin.H{
		//	"code": -1,
		//	"msg":  "group name cannot be empty",
		//})
		errReply(c, "group name cannot be empty")
		return
	}
	err := dao.JoinCommunity(value.(uint), name)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		errReply(c, err.Error())
		return
	}
	okReply(c)
}
