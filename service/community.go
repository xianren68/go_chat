package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_chat/common"
	"go_chat/dao"
	"go_chat/models"
	"net/http"
)

// NewGroup 新建群聊
func NewGroup(c *gin.Context) {
	value, _ := c.Get("userId")
	group := &models.Community{}
	err := c.BindJSON(group)
	if err != nil {
		zap.S().Warn("failed to read arguments ", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "请求参数错误",
		})
		return
	}
	// 群名为空
	if group.Name == "" {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "群名不能为空",
		})
		return
	}
	group.OwnerId = value.(uint)
	code := dao.CreateCommunity(group)
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

// GetGroupList 获取群列表
func GetGroupList(c *gin.Context) {
	// 从token中获取id
	value, _ := c.Get("userId")
	group, code := dao.GetCommunities(value.(uint))
	if code != 0 {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": -1,
		//	"msg":  err.Error(),
		//})
		common.ErrReply(c, code)
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
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "群名称不能为空",
		})
		return
	}
	code := dao.JoinCommunity(value.(uint), name)
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
