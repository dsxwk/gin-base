package utils

import "strings"

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

// InArrayString 检查字符串是否在切片中
// @param: str string, list []string
// @return: bool
func InArrayString(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
