package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gin-base/common/global"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
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

// @function: GetPageData
// @description: 获取分页数据
// @param: page, pageSize, total, data
// @return: global.PageData
func GetPageData(page int, pageSize int, total int64, data interface{}) global.PageData {
	return global.PageData{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		List:     data,
	}
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

// @function: BuildWhereClause
// @description: 构建查询条件
// @param: filters interface{}
// @param: tag 获取的标签 json、form 等
// @return: string, []interface{}
func BuildWhereClause(filters interface{}, tag string) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
		clause  string
	)

	val := reflect.ValueOf(filters)
	if val.Kind() != reflect.Struct {
		log.Fatalf("filters should be a struct, got %s", val.Kind())
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// 忽略零值字段
		if !value.IsZero() {
			// 使用 json 标签作为列名
			columnName := field.Tag.Get(tag)
			if columnName == "" {
				columnName = field.Name
			}

			// 根据字段类型构建查询条件
			switch value.Kind() {
			case reflect.String:
				clause = fmt.Sprintf("%s LIKE ?", columnName)
				args = append(args, "%"+value.String()+"%")
			default:
				clause = fmt.Sprintf("%s = ?", columnName)
				args = append(args, value.Interface())
			}

			clauses = append(clauses, clause)
		}
	}

	whereClause := strings.Join(clauses, " AND ")
	return whereClause, args
}

// @function: ToCamelCase
// @description: 将下划线分隔的字段名转换为驼峰命名
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

// @function: ConvertXStringToInt64Arr
// @description: *string 转 []int64
// @param: str *string
// @return: []int64
func ConvertXStringToInt64Arr(str *string) []int64 {
	var (
		data []int64
	)

	if str != nil {
		_ = json.Unmarshal([]byte(*str), &data)
		return data
	}

	return data
}

// @function: ConvertXStringToInt64Arr
// @description: *string 转 []string
// @param: str *string
// @return: []string
func ConvertXStringToStringArr(str *string) []string {
	var (
		data []string
	)

	if str != nil {
		_ = json.Unmarshal([]byte(*str), &data)
		return data
	}

	return data
}

// @function: RandomFileName
// @description: 生成随机文件名
// @param: ext string 文件后缀
// @return: string
func RandomFileName(ext string) string {
	// 获取当前时间的时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	// 组合时间戳和随机数生成文件名
	fileName := fmt.Sprintf("file_%d_%d."+ext, timestamp, randomNum)
	return fileName
}

// @function: RemoveDuplicates
// @description: 字符串切片去重
// @param: input []string
// @return: []string
func RemoveDuplicates(input []string) []string {
	var (
		outputData []string
	)

	uniqueDataMap := make(map[string]bool)

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
