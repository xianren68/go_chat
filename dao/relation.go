package dao

import (
	"go.uber.org/zap"
	"go_chat/global"
	"go_chat/models"
)

// FriendList 好友列表
func FriendList(userId uint) (friendList []*models.UserBasic) {
	relations := make([]*models.Relation, 0)
	tx := global.DB.Where("owner_id= ? and type= 1", userId).Find(&relations)
	if tx.RowsAffected == 0 {
		// 没有好友关系
		return
	}
	// 将所有的用户id提取出来
	userIds := make([]uint, len(relations))
	for i, relation := range relations {
		userIds[i] = relation.TargetId
	}
	global.DB.Where("id in ?", userIds).Find(&friendList)
	return
}

// AddFriend 添加好友
func AddFriend(p1, p2 uint) (code int) {
	// 同一个用户
	if p1 == p2 {
		code = 1009
		return
	}
	// 查询p1是否存在
	_, code = FindUserById(p1)
	if code != 0 {
		code = 1002
		return
	}
	// 查询p2是否存在
	_, code = FindUserById(p2)
	if code != 0 {
		code = 1002
		return
	}
	relation := &models.Relation{}
	// 判断好友信息是否存在
	tx := global.DB.Where("ownerid= ? and targetid= ? and type= 1", p1, p2).First(relation)
	if tx.RowsAffected == 1 {
		code = 1010
		return
	}
	tx = global.DB.Where("ownerid= ? and targetid= ? and type= 1", p2, p1).First(relation)
	if tx.RowsAffected == 1 {
		code = 1010
		return
	}
	// 开启事务
	tx = global.DB.Begin()
	// 添加好友（双向）
	relation.OwnerId = p1
	relation.TargetId = p2
	relation.Type = 1
	if t := tx.Create(relation); t.RowsAffected == 0 {
		code = 500
		zap.S().Warn("failed to create relation")
		// 事务回滚
		tx.Rollback()
	}
	relation = &models.Relation{}
	relation.OwnerId = p2
	relation.TargetId = p1
	relation.Type = 1
	if t := tx.Create(relation); t.RowsAffected == 0 {
		code = 500
		zap.S().Warn("failed to create relation")
		// 事务回滚
		tx.Rollback()
	}
	// 提交事务
	tx.Commit()
	return
}

// AddFriendByName 通过用户名添加好友
func AddFriendByName(id uint, name string) (code int) {
	// 判断用户名是否存在
	target := &models.UserBasic{}
	tx := global.DB.Where("name= ?", name).First(target)
	if tx.RowsAffected == 0 {
		code = 1002
		return
	}
	return AddFriend(id, target.ID)
}
