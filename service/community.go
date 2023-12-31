package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/common"
	"go_chat/dao"
	"go_chat/models"
	"net/http"
)

// NewGroup 新建群聊
func NewGroup(c *gin.Context) {
	value, _ := c.Get("userId")
	group := &models.Community{}
	_ = c.BindJSON(group)
	group.OwnerId = value.(uint)
	var code int
	if group.Name == "" {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "群名不能为空",
		})
		return
	}
	code = dao.CreateCommunity(group)
	if code != 0 {
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
	info := make(map[string]any)
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
		return
	}
	name := info["name"].(string)
	id := uint(info["id"].(float64))
	if name == "" {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "群名称不能为空",
		})
		return
	}
	code := dao.JoinCommunity(id, name)
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

// FindGroupByName 通过群名获取群信息
func FindGroupByName(c *gin.Context) {
	value := c.Query("name")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "failed to get name",
		})
		return
	}
	group, code := dao.FindGroupByName(value)
	if code != 0 {
		common.ErrReply(c, code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": group,
	})

}
