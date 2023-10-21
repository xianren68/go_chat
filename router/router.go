// Package router 路由信息
package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/middleware"
	"go_chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	// v1版本
	v1 := r.Group("v1")
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
	}
	// 关系相关
	relation := auth.Group("relation")
	{
		relation.GET("/friendlist", service.FriendList)
		relation.POST("/addfriend", service.AddFriend)
	}
	return r
}
