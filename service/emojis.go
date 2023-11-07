package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetEmoji 获取emoji表情
func GetEmoji(c *gin.Context) {
	emojis := []string{"😀", "😄", "😁", "😆", "😅", "🤣", "🙂", "🥰", "😍",
		"🤩", "😘", "🤑", "😋", "😜", "🤪", "🤗", "🤭", "🥵", "😡", "😈",
		"🤡", "👻", "🤖", "💦", "✌", "👍", "👎", "💗", "💤", "🤏", "👀", "🥷", "💃",
		"🐉", "🐳"}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": emojis,
		"msg":  "ok",
	})
}
