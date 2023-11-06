package models

// Message 消息结构体(表)
type Message struct {
	FromId   uint   `json:"from_id"`   // 发送者id
	TargetId uint   `json:"target_id"` // 接收者id
	GroupId  uint   `json:"group_id"`  // 群聊id
	Type     int    // 发送消息类型 （群发，私聊）
	Content  string // 消息内容
	SendTime int    // 发送时间
	Md       bool   // 是否是markdown
}
