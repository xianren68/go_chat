package models

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

// Message 消息结构体(表)
type Message struct {
	Model
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

// 全局channel
var upSendMsg = make(chan []byte, 1024)

// （接收客户端要发送的消息）
func recProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("failed to read message")
			return
		}
		// 将消息写入全局channel
		brodMsg(data)
	}
}

// 将消息写入全局channel
func brodMsg(data []byte) {
	upSendMsg <- data
}

func init() {
	go udpRecProc()
	go udpSendProc()
}

// 将消息写入udp服务器
func udpSendProc() {
	uC, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5000,
		Zone: "",
	})
	if err != nil {
		zap.S().Info("failed to connect udp")
		panic("failed to connect udp")
	}
	defer uC.Close()
	for bytes := range upSendMsg {
		_, err = uC.Write(bytes)
		if err != nil {
			zap.S().Info("failed to write udp")
			continue
		}
	}
}

// 开启udp服务
func udpRecProc() {
	uC, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5000,
		Zone: "",
	})
	if err != nil {
		zap.S().Info("failed to listen upd")
		panic("failed to listen upd")
	}
	// 获取客户端数据
	for {
		var buf = [1024]byte{}
		var n int
		n, err = uC.Read(buf[:])
		if err != nil {
			zap.S().Info("failed to get message")
			// 略过消息，继续下一条
			continue
		}
		// 解析消息
		disPatch(buf[:n])
	}
}

// 解析聊天
func disPatch(data []byte) {
	// 解析消息
	msg := &Message{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		zap.S().Info("failed to parse message")
		return
	}
	// 判断聊天类型
	switch msg.Type {
	// 私聊
	case 1:
		sendMsgAndSave(msg.TargetId, data)
	case 2:
		// sendGroup(msg.FormId, msg.TargetId)
	}

}
func sendMsgAndSave(id uint, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[id]
	rwLocker.Unlock()
	if !ok {
		zap.S().Info("user is not online")
		return
	}
	node.Data <- msg
}
