package middleware

import (
	"fmt"
	"gin-base/common/global"
	"gin-base/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Jwt struct {
}

var (
	Config = config.InitConfig()
)

func (s Jwt) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" || token == "null" {
			c.JSON(http.StatusOK, global.Response{global.Unauthorized, "请求未授权！", []string{}})
			c.Abort()
			return
		}

		data, err := s.Decode(token)
		if err != nil {
			c.JSON(http.StatusOK, global.Response{global.Unauthorized, err.Error(), []string{}})
			c.Abort()
			return
		}

		c.Set("user.id", data["id"])
	}
}

// 生成token
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

// 解析token
func (s Jwt) Decode(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("your_secret_key")
		return []byte(Config.Jwt.Key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
