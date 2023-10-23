package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

// Message 消息结构体(表)
type Message struct {
	gorm.Model
	FormId   uint   // 发送者id
	TargetId uint   // 接收者id
	Type     int    // 发送消息类型 （群发，私聊）
	Media    int    // 消息类型 （文字，图片，音频）
	Pic      string `json:"url"` // 图片相关
	Url      string // 文件相关
	Content  string // 消息内容
	Desc     string // 文件描述
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
	// 获取用户id
	userId, _ := c.Get("userId")
	Id := userId.(uint)
	// 将连接升级为socket
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
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
	for data := range node.Data {
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			zap.S().Info("failed to send message")
			return
		}
	}
}

// 接收消息
func recProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("failed to read message")
			return
		}
		// 解析消息
		msg := &Message{}
		err = json.Unmarshal(data, msg)
		if err != nil {
			zap.S().Info("failed to parse message")
		}
		// 判断消息类型
		switch msg.Type {
		// 私聊
		case 1:
			// 拿到接收人的node
			n, ok := clientMap[msg.TargetId]
			// 不在线/不存在
			if !ok {
				zap.S().Info("client is not exist")
				return
			}
			// 写入消息
			n.Data <- data
		}
	}
}
