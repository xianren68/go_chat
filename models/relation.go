package models

// Relation 关系结构体
type Relation struct {
	Model
	OwnerId  uint   // 谁的关系信息
	TargetId uint   // 与谁的关系
	Type     uint   // 关系类型 1.朋友 2.群组
	Desc     string // 描述
}

func (r *Relation) RelTableName() string {
	return "relation"
}
