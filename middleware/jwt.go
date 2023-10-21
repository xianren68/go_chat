package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

var tokenExpired = errors.New("token is expired")

// TOKEN_FLUSH token刷新时间
const TOKEN_FLUSH = 24

// 加密密钥
var jwtSecret = []byte("xianren")

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 使用id生成token
func GenerateToken(userId uint) (string, error) {
	// 设置token过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 2 * time.Hour)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go_chat",
		},
	}
	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token转化为字符串
	signedString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("generate token failed " + err.Error())
	}
	return signedString, nil

}

// ParseToken 解析token
func ParseToken(token string) (claims *Claims, err error) {
	claims = &Claims{}
	withClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		zap.S().Info("failed to parse token ", err)
		err = errors.New("failed to parse token")
		return
	}
	if myclaims, ok := withClaims.Claims.(*Claims); ok && withClaims.Valid {
		return myclaims, nil
	}
	err = errors.New("failed to parse token")
	return
}

// Jwt token中间件
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取token
		tokenHeader := ctx.GetHeader("Authorization")
		// 没有token
		if tokenHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token is null",
			})
			return
		}
		// 解析token格式
		tokenString := strings.SplitN(tokenHeader, " ", 2)
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": "-1",
				"msg":  "token format error",
			})
			return
		}
		// 解析token
		claims, err := ParseToken(tokenString[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		// 判断token是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token is expire",
			})
			return
		}
		// 判断token是否需要刷新
		if claims.ExpiresAt < time.Now().Add(time.Hour*TOKEN_FLUSH).Unix() {
			token, err := GenerateToken(claims.Id)
			// 成功了才刷新token
			if err == nil {
				ctx.Header("Authorization", "Bearer "+token)
			}
		}
		ctx.Set("userId", claims.Id)
	}
}
