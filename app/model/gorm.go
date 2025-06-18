package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gin-base/helper/utils"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type JsonTime time.Time

type JsonString []string

type JsonInt64 []int64

type BoolInt64 bool

type JsonMap map[string]interface{}

// Search struct公共搜索
// 注：会过滤零值
// 用例：global.DB.Scopes(model.Search(search)).Find(&models)
// 示例值：
// UserSearch 用户搜索
//
//	type UserSearch struct {
//		Username string `form:"username" label:"用户名"`
//		FullName string `form:"fullName" label:"姓名"`
//		Nickname string `form:"nickname" label:"昵称"`
//		Gender   int    `form:"gender" label:"性别"`
//	}
func Search(filters interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		val := reflect.ValueOf(filters)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			fieldStruct := typ.Field(i)
			value := val.Field(i)
			if value.IsZero() {
				continue
			}
			// 优先 form
			tag := fieldStruct.Tag.Get("form")
			if tag == "" {
				tag = fieldStruct.Tag.Get("json")
			}
			tag = utils.CamelToSnake(tag)
			if strings.Contains(tag, ".") {
				parts := strings.Split(tag, ".")
				baseField := parts[0]
				jsonPath := "$." + strings.Join(parts[1:], ".")
				db = db.Where(baseField+"->>'"+jsonPath+"' LIKE ?", value.Interface())
			} else {
				db = db.Where(tag+" LIKE ?", value.Interface())
			}
		}
		return db
	}
}

// SearchMap map公共搜索
// 用例：global.DB.Scopes(model.SearchMap(search)).Find(&models)
// 示例值: 如前端传值到 __search{} 中, __search示例：
// operator: >,<, >=, <=, =, !=, in, not in, between, not between, is null, is not null, like, not like, left like, right like
//
//	__search: {
//	   "field": {"operator": ">", "value": 10},
//	   "field": {"operator": "<", "value": 10},
//	   "field": {"operator": ">=", "value": 10},
//	   "field": {"operator": "<=", "value": 10},
//	   "field": {"operator": "=", "value": 10},
//	   "field": {"operator": "!=", "value": 10},
//	   "field": {"operator": "in", "value": [1,2,3]},
//	   "field": {"operator": "not in", "value": [1,2,3]},
//	   "field": {"operator": "between", "value": [1,2]},
//	   "field": {"operator": "not between", "value": [1,2]},
//	   "field": {"operator": "is null"},
//	   "field": {"operator": "is not null"},
//	   "field": {"operator": "like", "value": "admin"},
//	   "field": {"operator": "left like", "value": "test"},
//	   "field": {"operator": "right like", "value": "138"},
//	   "jsonField.field": {"operator": ">", "value": 10},
//	   ...
//	}
func SearchMap(filters map[string]interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key, raw := range filters {
			var (
				op    = "="
				valIf interface{}
			)
			if m, ok := raw.(map[string]interface{}); ok {
				if v, ok := m["operator"].(string); ok {
					op = strings.ToLower(v)
				}
				valIf = m["value"]
			} else {
				valIf = raw
			}

			var (
				field string
			)
			if strings.Contains(key, ".") {
				parts := strings.Split(key, ".")
				baseField := parts[0]
				jsonPath := "$." + strings.Join(parts[1:], ".")
				field = fmt.Sprintf("%s->>'%s'", baseField, jsonPath)
			} else {
				field = key
			}

			switch op {
			case "=":
				db = db.Where(fmt.Sprintf("%s = ?", field), valIf)
			case "!=", "<>":
				db = db.Where(fmt.Sprintf("%s <> ?", field), valIf)
			case ">":
				db = db.Where(fmt.Sprintf("%s > ?", field), valIf)
			case "<":
				db = db.Where(fmt.Sprintf("%s < ?", field), valIf)
			case ">=":
				db = db.Where(fmt.Sprintf("%s >= ?", field), valIf)
			case "<=":
				db = db.Where(fmt.Sprintf("%s <= ?", field), valIf)
			case "in":
				db = db.Where(fmt.Sprintf("%s IN (?)", field), valIf)
			case "not in":
				db = db.Where(fmt.Sprintf("%s NOT IN (?)", field), valIf)
			case "like":
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+fmt.Sprint(valIf)+"%")
			case "left like":
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprint(valIf)+"%")
			case "right like":
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+fmt.Sprint(valIf))
			case "not like":
				db = db.Where(fmt.Sprintf("%s NOT LIKE ?", field), "%"+fmt.Sprint(valIf)+"%")
			case "not left like":
				db = db.Where(fmt.Sprintf("%s NOT LIKE ?", field), fmt.Sprint(valIf)+"%")
			case "not right like":
				db = db.Where(fmt.Sprintf("%s NOT LIKE ?", field), "%"+fmt.Sprint(valIf))
			case "between":
				if vals, ok := valIf.([]interface{}); ok && len(vals) == 2 {
					db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), vals[0], vals[1])
				}
			case "not between":
				if vals, ok := valIf.([]interface{}); ok && len(vals) == 2 {
					db = db.Where(fmt.Sprintf("%s NOT BETWEEN ? AND ?", field), vals[0], vals[1])
				}
			case "is null":
				db = db.Where(fmt.Sprintf("%s IS NULL", field))
			case "is not null":
				db = db.Where(fmt.Sprintf("%s IS NOT NULL", field))
			default:
				db = db.Where(fmt.Sprintf("%s %s ?", field, op), valIf)
			}
		}
		return db
	}
}

// MarshalJSON 模型时间格式化公共方法
func (t *JsonTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`""`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JsonTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*t = JsonTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = JsonTime(v)
		return nil
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
		return nil
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
		return nil
	default:
		return fmt.Errorf("cannot convert %v to timestamp", value)
	}
}

func (j JsonString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonString) Scan(value interface{}) error {
	if value == nil {
		*j = JsonString{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonString", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j JsonInt64) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonInt64) Scan(value interface{}) error {
	if value == nil {
		*j = JsonInt64{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonInt64", value)
	}
	return json.Unmarshal(bytes, j)
}

// Value 写入数据库时：true → 1，false → 2
func (b BoolInt64) Value() (driver.Value, error) {
	if b {
		return 1, nil
	}
	return 2, nil
}

// Scan 读取数据库时：1 → true，2 → false（默认 false）
func (b *BoolInt64) Scan(value interface{}) error {
	if value == nil {
		*b = false
		return nil
	}
	var v int64
	switch val := value.(type) {
	case int64:
		v = val
	case int:
		v = int64(val)
	case int32:
		v = int64(val)
	case uint8:
		v = int64(val)
	case []byte:
		i, err := strconv.Atoi(string(val))
		if err != nil {
			return err
		}
		v = int64(i)
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		v = int64(i)
	default:
		return fmt.Errorf("unsupported Scan type for BoolInt12: %T", value)
	}

	*b = v == 1
	return nil
}

func (j *JsonMap) Scan(value interface{}) error {
	if value == nil {
		*j = make(JsonMap)
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("JsonMap: Scan source is not []byte")
	}
	return json.Unmarshal(bytes, j)
}

func (j JsonMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// JsonTime
func (t *JsonTime) String() string {
	if t == nil {
		return ""
	}
	return time.Time(*t).Format("2006-01-02 15:04:05")
}

// JsonString
func (j JsonString) String() string {
	b, err := json.Marshal(j)
	if err != nil {
		return "[]"
	}
	return string(b)
}

// JsonInt64
func (j JsonInt64) String() string {
	b, err := json.Marshal(j)
	if err != nil {
		return "[]"
	}
	return string(b)
}

// BoolInt64
func (b BoolInt64) String() string {
	if b {
		return "true"
	}
	return "false"
}

// JsonMap
func (j JsonMap) String() string {
	b, err := json.Marshal(j)
	if err != nil {
		return "{}"
	}
	return string(b)
}
