package middleware

import (
	"gin-base/common/base"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
)

type I18n struct {
	base.BaseMiddleware
}

// Handle 多语言翻译
func (s I18n) Handle() gin.HandlerFunc {
	var (
		lang string
	)

	return func(c *gin.Context) {
		lang = c.GetHeader("Accept-Language")
		if strutil.StartsWith(lang, "en") {
			lang = "en"
		} else {
			lang = "zh"
		}

		c.Set("lang", lang)

		c.Next()
	}
}
