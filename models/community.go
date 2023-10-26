package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-" `
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Community 群聊结构体
type Community struct {
	Model
	Name    string // 群名称
	OwnerId uint   // 群主
	Type    int    // 群类型
	Image   string // 头像
	Desc    string // 描述
}
