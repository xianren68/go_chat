package models

import "gorm.io/gorm"

// Community 群聊结构体
type Community struct {
	gorm.Model
	Name    string // 群名称
	OwnerId uint   // 群主
	Type    int    // 群类型
	Image   string // 头像
	Desc    string // 描述
}
