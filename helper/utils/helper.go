package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// @function: FormatTime
// @description: 格式化时间
// @param: t string
// @return: string
func FormatTime(t string) string {
	ts, err := time.Parse(time.RFC3339, t)

	if err != nil {
		panic(err)
	}

	return ts.Format("2006-01-02 15:04:05")
}

// @function: Pagination
// @description: 计算分页
// @param: page, pageSize int
// @return: int, int
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

// @function: BcryptHash
// @description: 使用bcrypt对密码进行加密
// @param: password string
// @return: string
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes)
}

// @function: BcryptCheck
// @description: 对比明文密码和数据库的哈希值
// @param: password, hash string
// @return: bool
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// @function: Md5
// @description: md5加密
// @param: str []byte
// @return: string
func Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)

	return hex.EncodeToString(h.Sum(b))
}

// @function: KeyIsExist
// @description: 检查map键名是否存在
// @param: data map[string]interface{}, key string
// @return: bool
func KeyIsExist(data map[string]interface{}, key string) bool {
	_, exists := data[key]

	return exists
}

// @function: CamelToSnake
// @description: 驼峰转下划线
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

// @function: CamelToSnakeMap
// @description: 驼峰键名数据转下划线键名数据
// @param: data map[string]interface{}
// @return: map[string]interface{}
func CamelToSnakeMap(data map[string]interface{}) map[string]interface{} {
	returnData := make(map[string]interface{})

	for k, v := range data {
		returnData[CamelToSnake(k)] = v
	}

	return returnData
}

// @function: FilterStructToMap
// @description: 根据结构体过滤请求字段
// @param: req map[string]interface{}, filterStruct interface{}
// @return: result map[string]interface{}
func FilterStructToMap(req map[string]interface{}, filterStruct interface{}) (result map[string]interface{}) {
	req = make(map[string]interface{})

	// 使用反射遍历结构体字段,只获取在结构体中定义的字段
	value := reflect.ValueOf(filterStruct)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		column := field.Tag.Get("json")

		if KeyIsExist(req, column) {
			result[column] = req[column]
		}
	}

	return result
}
