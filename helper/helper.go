package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"
)

// FormatTime 泛型方法，支持 *time.Time、time.Time、int/int64（unix秒）、string（时间字符串）、float64（unix秒）
// @param: t
// @return: string
func FormatTime[T any](t T) string {
	switch v := any(t).(type) {
	case *time.Time:
		if v == nil {
			return ""
		}
		return v.Format("2006-01-02 15:04:05")
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case int64:
		return time.Unix(v, 0).Format("2006-01-02 15:04:05")
	case int:
		return time.Unix(int64(v), 0).Format("2006-01-02 15:04:05")
	case float64:
		return time.Unix(int64(v), 0).Format("2006-01-02 15:04:05")
	case string:
		layouts := []string{
			"2006-01-02 15:04:05",
			"2006-01-02",
			time.RFC3339,
		}
		for _, layout := range layouts {
			if tm, err := time.Parse(layout, v); err == nil {
				return tm.Format("2006-01-02 15:04:05")
			}
		}
		if unix, err := strconv.ParseInt(v, 10, 64); err == nil {
			return time.Unix(unix, 0).Format("2006-01-02 15:04:05")
		}
	}

	return ""
}

// Pagination 计算分页
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

// BcryptHash 使用bcrypt对密码进行加密
// @param: password string
// @return: string
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
// @param: password, hash string
// @return: bool
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// Md5 md5加密
// @param: str []byte, b ...byte
// @return: string
func Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)

	return hex.EncodeToString(h.Sum(b))
}

// HasKey 检查map键名是否存在
// @param: data map[string]interface{}, key string
// @return: bool
func HasKey(data map[string]interface{}, key string) bool {
	_, exists := data[key]

	return exists
}

// CamelToSnakeMap 驼峰键名数据转下划线键名数据
// @param: data map[string]interface{}
// @return: map[string]interface{}
func CamelToSnakeMap(data map[string]interface{}) map[string]interface{} {
	returnData := make(map[string]interface{})

	for k, v := range data {
		returnData[CamelToSnake(k)] = v
	}

	return returnData
}

// ToCamelCase 将下划线分隔的字段名转换为驼峰命名
// @param: s string
// @return: string
func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		if i > 0 {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

// ConvertToArr 泛型方法 *string 转 []T
// 示例1 arr1 := ConvertToArr[int64](str)    // *string 转 []int64
// 示例2 arr2 := ConvertToArr[string](str)   // *string 转 []string
func ConvertToArr[T any](str *string) []T {
	var data []T
	if str != nil {
		_ = json.Unmarshal([]byte(*str), &data)
	}
	return data
}

// ArrayColumn 泛型 提取结构体切片的某字段
//
//	示例 type MenuRoles struct {
//	   Name string
//	   ID   int64
//	}
//
// names := ArrayColumn(MenuRoles, func(m *MenuRoles) string { return m.Name }) // []string
// ids := ArrayColumn(MenuRoles, func(m *MenuRoles) int64 { return m.ID })      // []int64
func ArrayColumn[T any, R any](arr []*T, getter func(*T) R) []R {
	result := make([]R, 0, len(arr))
	for _, v := range arr {
		result = append(result, getter(v))
	}
	return result
}

// RandomFilename 生成随机文件名
// @param: ext string 文件后缀
// @return: string
func RandomFilename(ext string) string {
	// 获取当前时间的时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	// 组合时间戳和随机数生成文件名
	fileName := fmt.Sprintf("file_%d_%d."+ext, timestamp, randomNum)

	return fileName
}

// RemoveDuplicates 泛型切片去重
// @param: input []T
// @return: []T
func RemoveDuplicates[T comparable](input []T) []T {
	var outputData []T
	uniqueDataMap := make(map[T]bool)
	for _, value := range input {
		if !uniqueDataMap[value] {
			uniqueDataMap[value] = true
			outputData = append(outputData, value)
		}
	}
	return outputData
}

// BatchUpdateSql 生成批量更新 SQL 语句
// table: 数据表名称
// data: 更新数据，每个元素是一个包含列名和值的映射
// primaryKey: 主键字段名
// conditions: 附加的 WHERE 条件（可选）
// 返回值是生成的 SQL 语句和对应的参数值
// 调用示例: sql, values := BatchUpdateSql("table", data, "id", map[string]interface{}{"status": 1})
// global.DB.Exec(sql, values...)
func BatchUpdateSql(table string, data []map[string]interface{}, primaryKey string, conditions map[string]interface{}) (string, []interface{}) {
	if len(data) == 0 || primaryKey == "" {
		return "", nil
	}

	var (
		setParts []string
		values   []interface{}
		idValues []interface{}
	)

	// 确保处理所有的列名
	columns := make(map[string]bool)
	for _, row := range data {
		for column := range row {
			if column != primaryKey {
				columns[column] = true
			}
		}
		// 将 id 值转换为正确的类型
		idValues = append(idValues, row[primaryKey])
	}

	// 生成 SET 子句
	for column := range columns {
		caseClause := fmt.Sprintf("`%s` = CASE `%s` ", column, primaryKey)
		for _, row := range data {
			caseClause += "WHEN ? THEN ? "
			values = append(values, row[primaryKey], row[column])
		}
		caseClause += "END"
		setParts = append(setParts, caseClause)
	}

	setClause := strings.Join(setParts, ", ")

	var idValuesStr []string
	for _, v := range idValues {
		idValuesStr = append(idValuesStr, fmt.Sprintf("%v", v))
	}
	// 生成 WHERE 子句
	whereParts := []string{fmt.Sprintf("`%s` IN (%s)", primaryKey, strings.Join(idValuesStr, ","))}
	for k, v := range conditions {
		whereParts = append(whereParts, fmt.Sprintf("`%s` = ?", k))
		values = append(values, v)
	}

	whereClause := strings.Join(whereParts, " AND deleted_at IS NULL")

	sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", table, setClause, whereClause)

	return sql, values
}

// StructToMapFilter 将结构体转换为map并过滤指定的字段
func StructToMapFilter(data interface{}, excludeFields []string) map[string]interface{} {
	s := structs.New(data)
	s.TagName = "json"
	val := s.Map()

	res := make(map[string]interface{})
	for k, v := range val {
		if slices.Contains(excludeFields, k) {
			continue
		}

		res[CamelToSnake(k)] = v
	}

	return res
}
