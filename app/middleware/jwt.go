package middleware

import (
	"fmt"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Jwt struct {
	common.BaseMiddleware
}

var (
	Config = config.InitConfig()
)

// JwtMiddleware jwt中间件
func (s Jwt) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" || token == "null" {
			global.Log.Error(global.Unauthorized.Message)
			global.ApiResponse(c, global.Unauthorized)
			return
		}

		data, err := s.Decode(token)
		if err != nil {
			global.Log.Error(global.Unauthorized.Message)
			global.ApiResponse(c, global.Unauthorized, err.Error())
			return
		}

		c.Set("user.id", data["id"])
	}
}

// Encode 生成token
func (s Jwt) Encode(id int64, exp int64) (string, int64, error) {
	if exp == 0 {
		exp = time.Now().Add(time.Duration(Config.Jwt.Exp) * time.Second).Unix()
	} else {
		exp = time.Now().Add(time.Duration(exp) * time.Second).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": exp,
	})

	jwtToken, err := token.SignedString([]byte(Config.Jwt.Key))

	return jwtToken, exp, err
}

// Decode 解析token
func (s Jwt) Decode(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("your_secret_key")
		return []byte(Config.Jwt.Key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token 解析失败: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的 token")
}
