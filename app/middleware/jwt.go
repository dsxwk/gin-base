package middleware

import (
	"fmt"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Jwt struct {
	base.BaseMiddleware
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
func (s Jwt) Encode(id int64, accessExp int64) (string, int64, error) {
	var (
		now = time.Now()
	)

	if accessExp == 0 {
		accessExp = now.Add(time.Duration(Config.Jwt.Exp) * time.Second).Unix()
	} else {
		accessExp = now.Add(time.Duration(accessExp) * time.Second).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": accessExp,
	})

	jwtToken, err := token.SignedString([]byte(Config.Jwt.Key))

	return jwtToken, accessExp, err
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

// WithRefresh 刷新token
func (s Jwt) WithRefresh(id, accessExp, refreshExp int64) (accessToken, refreshToken string, tExp, rExp int64, err error) {
	var (
		now = time.Now()
	)

	if accessExp == 0 {
		accessExp = now.Add(time.Duration(Config.Jwt.Exp) * time.Second).Unix()
	} else {
		accessExp = now.Add(time.Duration(accessExp) * time.Second).Unix()
	}

	if refreshExp == 0 {
		refreshExp = now.Add(time.Duration(Config.Jwt.RefreshExp) * time.Second).Unix()
	} else {
		refreshExp = now.Add(time.Duration(refreshExp) * time.Second).Unix()
	}

	// Access Token
	accessClaims := jwt.MapClaims{
		"id":  id,
		"exp": accessExp,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString([]byte(Config.Jwt.Key))
	if err != nil {
		return "", "", 0, 0, err
	}

	// Refresh Token
	refreshClaims := jwt.MapClaims{
		"id":  id,
		"exp": refreshExp,
		"typ": "refresh", // 标记为 refresh token
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString([]byte(Config.Jwt.Key))
	if err != nil {
		return "", "", 0, 0, err
	}

	return accessToken, refreshToken, accessExp, refreshExp, nil
}
