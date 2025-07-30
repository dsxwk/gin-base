package tests

import (
	"gin-base/common/extend/i18n"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 翻译
func TestI18n_T(t *testing.T) {
	i18n.Init(true)

	t.Run("zh translation", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("lang", "zh")

		msg := i18n.T(c, "login.code", nil)
		assert.Equal(t, "验证码", msg)
	})

	t.Run("zh translation", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("lang", "zh")

		msg := i18n.T(c, "login.success", map[string]interface{}{
			"name": "admin",
		})
		assert.Equal(t, "admin,登录成功", msg)
	})

	t.Run("en translation", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("lang", "en")

		msg := i18n.T(c, "login.code", nil)
		assert.Equal(t, "Code", msg)
	})

	t.Run("en translation", func(t *testing.T) {
		c := &gin.Context{}
		c.Set("lang", "en")

		msg := i18n.T(c, "login.success", map[string]interface{}{
			"name": "admin",
		})
		assert.Equal(t, "admin,Login Success", msg)
	})
}
