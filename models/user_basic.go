// Package models 数据库模型
package models

import (
	"gorm.io/gorm"
	"time"
)

// UserBasic 用户结构体
type UserBasic struct {
	gorm.Model    // 包含唯一id，更新/删除时间等
	Name          string
	PassWord      string
	Avatar        string
	Gender        string `gorm:"column:gender;default:male;type:varchar(6)"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string     `valid:"ipv4" json:"-"`
	ClientPort    string     `json:"-"`
	Salt          string     `json:"-"` // 加密盐值
	LoginTime     *time.Time `gorm:"column:login_time" json:"-"`
	HeartBeatTime *time.Time `gorm:"column:heart_beat_time" json:"-"`
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"-"`
	IsLoginOut    bool       `json:"-"`
	DeviceInfo    string     `json:"-"` //登录设备
}

// UserTableName 返回表名称
func (table *UserBasic) UserTableName() string {
	return "user_basic"
}
