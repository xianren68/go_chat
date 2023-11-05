// Package router 路由信息
package router

import (
	"go_chat/middleware"
	"go_chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 跨域
	r.Use(middleware.Cors())
	// v1版本
	v1 := r.Group("v1")
	v1.GET("/chat", service.SendUserMsg)
	v1.POST("/login", service.Login)
	v1.POST("/new", service.NewUser)
	// 鉴权
	auth := v1.Group("auth")
	auth.Use(middleware.Jwt())
	// 用户相关
	user := auth.Group("user")
	{
		user.GET("/list", service.List)
		user.DELETE("/delete", service.DeleteUser)
		user.PUT("/update", service.UpdateUser)
		user.GET("/unread", service.GetUnread)
		user.GET("/emojis", service.GetEmoji)
	}
	// 关系相关
	relation := auth.Group("relation")
	{
		relation.GET("/friendlist", service.FriendList)
		relation.POST("/addfriend", service.AddFriend)
		relation.GET("/grouplist", service.GetGroupList)
		relation.POST("/newgroup", service.NewGroup)
		relation.POST("joingroup", service.JoinGroup)
	}
	return r
}
