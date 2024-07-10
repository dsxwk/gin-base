package utils

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 格式化时间
func FormatTime(t string) string {
	ts, err := time.Parse(time.RFC3339, t)

	if err != nil {
		panic(err)
	}

	return ts.Format("2006-01-02 15:04:05")
}

// 计算分页
func Pagination(page, pageSize int) (int, int) {
	// 获取分页默认为第一页，每页10条记录
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	limit := pageSize

	return offset, limit
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
