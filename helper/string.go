package helper

import (
	"regexp"
	"strings"
)

// FirstUpper 首字母大写
// @param: s string
// @return: string
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 首字母小写
// @param: s string
// @return: string
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// Snake 字符串转驼峰
// @param: s string
// @return: string
func Snake(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// CamelToSnake 驼峰转下划线
// @param: field string
// @return: string
func CamelToSnake(field string) string {
	// 使用正则表达式匹配大写字母
	reg := regexp.MustCompile("[A-Z]")

	// 将匹配到的大写字母替换为下划线+小写字母
	snake := reg.ReplaceAllStringFunc(field, func(match string) string {
		return "_" + strings.ToLower(match)
	})

	// 去掉开头可能的下划线
	snake = strings.TrimPrefix(snake, "_")

	return snake
}
