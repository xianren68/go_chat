package models

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"go_chat/common"
	"go_chat/global"
	"go_chat/middleware"
	"gopkg.in/fatih/set.v0"
	"net/http"
	"strconv"
	"sync"
)

// Message 消息结构体(表)
type Message struct {
	FromId   uint   `json:"from_id"`   // 发送者id
	TargetId uint   `json:"target_id"` // 接收者id
	groupId  uint   `json:"group_id"`  // 群聊id
	Type     int    // 发送消息类型 （群发，私聊）
	Content  string // 消息内容
}

// Node 构造连接
type Node struct {
	Conn      *websocket.Conn // socket连接
	Addr      string          // 客户端地址
	Data      chan []byte     // 消息内容
	GroupSets set.Interface
}

// 所有的socket连接
var clientMap = make(map[uint]*Node)

// 读写锁 绑定node时需要线程安全
var rwLocker sync.RWMutex

// Chat 升级连接，初始化node
func Chat(c *gin.Context) {
	// 获取token
	token := c.Request.Header.Get("Sec-WebSocket-Protocol")
	if token == "" {
		common.ErrReply(
			c,
			1005,
		)
		c.Abort()
		return
	}
	// 解析token
	claims, code := middleware.ParseToken(token)
	if code != 0 {
		common.ErrReply(c, code)
		c.Abort()
		return
	}
	// 获取id
	Id := claims.Id
	// 将连接升级为socket
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-Websocket-Protocol")},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Warn(err)
		return
	}
	// 获取socket连接，建立消息节点
	node := &Node{
		Conn:      conn,
		Data:      make(chan []byte, 100),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 绑定用户id和socket连接
	rwLocker.Lock()
	clientMap[Id] = node
	rwLocker.Unlock()

	// 发送消息
	go sendProc(node)
	// 接收消息
	go recProc(node)
}

// 发送消息
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.Data:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("failed to send message")
				return
			}
		}
	}
}

// （接收客户端要发送的消息）
func recProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("failed to read message", err)
			return
		}
		disPatch(data)

	}
}

// 解析聊天
func disPatch(data []byte) {
	// 解析消息
	msg := &Message{}
	err := json.Unmarshal(data, msg)

	if err != nil {
		zap.S().Info("failed to parse message", err)
		return
	}

	// 判断聊天类型
	switch msg.Type {
	// 私聊
	case 1:
		sendMsgAndSave(msg.TargetId, data)
		// 群聊
	case 2:
		// sendGroup(msg.FormId, msg.TargetId)
	case 3:
		// 请求添加好友
	}

}
func sendMsgAndSave(id uint, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[id]
	rwLocker.Unlock()
	// 用户不在线，存进redis缓存中
	if !ok {
		saveRedis(id, msg)
		zap.S().Info("user is not online")
		return
	}
	node.Data <- msg
}

func saveRedis(id uint, msg []byte) {
	ctx := context.Background()
	key := "list" + strconv.Itoa(int(id))
	_, _ = global.RedisDB.LPush(ctx, key, string(msg)).Result()

}
