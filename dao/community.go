package dao

import (
	"errors"
	"go_chat/global"
	"go_chat/models"
)

// FindUsers 返回群组中用户id
func FindUsers(groupId uint) (group []uint, err error) {
	relations := make([]*models.Relation, 0)
	if tx := global.DB.Where("target_id= ? and type = 2", groupId).Find(&relations); tx.RowsAffected == 0 {
		err = errors.New("no group member information found")
		return
	}
	// 提取群员id
	for _, relation := range relations {
		group = append(group, relation.OwnerId)
	}
	return
}

// CreateCommunity 新建群聊
func CreateCommunity(community *models.Community) (err error) {
	// 查询是否已经存在此群聊
	if tx := global.DB.Where("name= ?", community.Name).First(community); tx.RowsAffected == 1 {
		err = errors.New("group chat already exists")
	}
	// 开启事务
	tx := global.DB.Begin()
	// 创建群聊
	if t := tx.Create(community); t.RowsAffected == 0 {
		err = errors.New("failed to create group")
		// 事务回滚
		tx.Rollback()
		return
	}
	// 创建关系连接
	relation := &models.Relation{}
	relation.Type = 2
	relation.TargetId = community.ID
	relation.OwnerId = community.OwnerId
	if t := tx.Create(relation); t.RowsAffected == 0 {
		err = errors.New("failed to create group")
		// 事务回滚
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// GetCommunities 获取群列表
func GetCommunities(ownerId uint) (group []*models.Community, err error) {
	relations := make([]*models.Relation, 0)
	// 判断用户是否有加群
	if tx := global.DB.Where("owner_id= ? and type=2", ownerId).Find(&relations); tx.RowsAffected == 0 {
		err = errors.New("user dose not have any groups")
		return
	}
	communityId := make([]uint, len(relations))
	for i, relation := range relations {
		communityId[i] = relation.TargetId
	}
	// 获取所有的群聊信息
	if tx := global.DB.Where("id in ?", communityId).Find(&group); tx.RowsAffected == 0 {
		err = errors.New("failed to get group information")
		return
	}
	return
}

// JoinCommunity 加入群聊
func JoinCommunity(userId uint, groupName string) (err error) {
	community := &models.Community{}
	// 搜索群聊是否存在
	if tx := global.DB.Where("name= ?", groupName).First(community); tx.RowsAffected == 0 {
		err = errors.New("this group is not exist")
		return
	}
	// 判断是否已经加入过
	tx := global.DB.Where("owner_id= ? and target_id= ? and type=2", userId, community.ID).First(&models.Relation{})
	if tx.RowsAffected == 1 {
		err = errors.New("already joined this group")
		return
	}
	// 加入群聊
	relation := &models.Relation{}
	relation.OwnerId = userId
	relation.TargetId = community.ID
	relation.Type = 2
	if tx = global.DB.Create(relation); tx.RowsAffected == 0 {
		err = errors.New("failed to join group")
		return
	}
	return
}
