package dao

import (
	"errors"
	"go.uber.org/zap"
	"go_chat/global"
	"go_chat/models"
	"strconv"
)

// FriendList 好友列表
func FriendList(userId uint) (friendList []*models.UserBasic, err error) {
	relations := make([]*models.Relation, 0)
	tx := global.DB.Where("owner_id= ? and type= 1", userId).Find(&relations)
	if tx.RowsAffected == 0 {
		err = errors.New(strconv.Itoa(int(userId)) + " friendList is not found")
		return
	}
	// 将所有的用户id提取出来
	userIds := make([]uint, len(relations))
	for i, relation := range relations {
		userIds[i] = relation.TargetId
	}
	find := global.DB.Where("id in ?", userIds).Find(&friendList)
	if find.RowsAffected == 0 {
		err = errors.New(strconv.Itoa(int(userId)) + " friendList is not found")
		return
	}
	return
}

// AddFriend 添加好友
func AddFriend(p1, p2 uint) (code int, err error) {
	// 同一个用户
	if p1 == p2 {
		code = -2
		err = errors.New("it is the same")
		return
	}
	// 查询p1是否存在
	_, err = FindUserById(p1)
	if err != nil {
		code = -1
		err = errors.New("user" + strconv.Itoa(int(p1)) + "is not exist")
		return
	}
	// 查询p2是否存在
	_, err = FindUserById(p2)
	if err != nil {
		code = -1
		err = errors.New("user" + strconv.Itoa(int(p1)) + "is not exist")
		return
	}
	relation := &models.Relation{}
	// 判断好友信息是否存在
	tx := global.DB.Where("ownerid= ? and targetid= ? and type= 1", p1, p2).First(relation)
	if tx.RowsAffected == 1 {
		code = 0
		err = errors.New("this relation is exist")
		return
	}
	tx = global.DB.Where("ownerid= ? and targetid= ? and type= 1", p2, p1).First(relation)
	if tx.RowsAffected == 1 {
		code = 0
		err = errors.New("this relation is exist")
		return
	}
	// 开启事务
	tx = global.DB.Begin()
	// 添加好友（双向）
	relation.OwnerId = p1
	relation.TargetId = p2
	relation.Type = 1
	if t := tx.Create(relation); t.RowsAffected == 0 {
		code = -1
		zap.S().Warn("failed to create relation")
		// 事务回滚
		tx.Rollback()
		err = errors.New("failed to create relation")
	}
	relation = &models.Relation{}
	relation.OwnerId = p2
	relation.TargetId = p1
	relation.Type = 1
	if t := tx.Create(relation); t.RowsAffected == 0 {
		code = -1
		zap.S().Warn("failed to create relation")
		// 事务回滚
		tx.Rollback()
		err = errors.New("failed to create relation")
	}
	// 提交事务
	tx.Commit()
	code = 1
	return
}

// AddFriendByName 通过用户名添加好友
func AddFriendByName(id uint, name string) (code int, err error) {
	// 判断用户名是否存在
	target := &models.UserBasic{}
	tx := global.DB.Where("name= ?", name).First(target)
	if tx.RowsAffected == 0 {
		code = -1
		err = errors.New(name + "is not exist")
		return
	}
	return AddFriend(id, target.ID)
}
