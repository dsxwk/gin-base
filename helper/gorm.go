package helper

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// gorm operator
var operator = map[string]bool{
	"or": true, ">": true, "<": true, ">=": true, "<=": true,
	"!=": true, "<>": true, "in": true, "not in": true, "like": true,
	"not like": true, "between": true, "not between": true,
	"is null": true, "is not null": true, "exists": true, "not exists": true,
	"in subquery": true, "not in subquery": true,
	"json contains": true, "json not contains": true,
	"json contains key": true, "json not contains key": true,
	"json contains path": true, "json not contains path": true,
	"json contains path value": true, "json not contains path value": true,
	"json contains path key": true, "json not contains path key": true,
	"json contains path key value": true, "json not contains path key value": true,
}

// isOperator 判断是否是操作符
func isOperator(key string) bool {
	_, ok := operator[strings.ToLower(key)]
	return ok
}

// ApplyFilters 递归解析 __filters
func ApplyFilters(db *gorm.DB, conditions map[string]interface{}, useOr bool) *gorm.DB {
	if useOr {
		// OR 包装
		return db.Or(func(tx *gorm.DB) *gorm.DB {
			return parseConditions(tx, conditions)
		})
	}
	return parseConditions(db, conditions)
}

// parseConditions 解析条件
func parseConditions(db *gorm.DB, conditions map[string]interface{}) *gorm.DB {
	for key, val := range conditions {
		keyLower := strings.ToLower(key)
		if isOperator(keyLower) {
			// 操作符
			switch keyLower {
			case "or":
				// or: [ {...}, {...} ]
				if arr, ok := val.([]interface{}); ok {
					for _, item := range arr {
						if m, ok := item.(map[string]interface{}); ok {
							db = db.Or(func(tx *gorm.DB) *gorm.DB {
								return parseConditions(tx, m)
							})
						}
					}
				}
			case "and":
				// and
				if arr, ok := val.([]interface{}); ok {
					for _, item := range arr {
						if m, ok := item.(map[string]interface{}); ok {
							db = db.Where(func(tx *gorm.DB) *gorm.DB {
								return parseConditions(tx, m)
							})
						}
					}
				}
			default:
				// 其他操作符
			}
		} else {
			// 非操作符，视作字段名，值可能是：
			switch v := val.(type) {
			case map[string]interface{}:
				// 可能 { field: { op: value } } 形式
				for opKey, opVal := range v {
					opKeyLower := strings.ToLower(opKey)
					if isOperator(opKeyLower) {
						db = applyCondition(db, key, opKeyLower, opVal)
					} else {
						// 递归
						if m, ok := opVal.(map[string]interface{}); ok {
							db = ApplyFilters(db, map[string]interface{}{opKey: m}, false)
						}
					}
				}
			case []interface{}:
				// 数组，默认当in处理
				db = applyCondition(db, key, "in", v)
			default:
				// 基础值，默认等于
				db = applyCondition(db, key, "=", v)
			}
		}
	}
	return db
}

// applyCondition 应用单条条件，兼容 preload 和 JSON
func applyCondition(db *gorm.DB, field, op string, value interface{}) *gorm.DB {
	op = strings.ToLower(op)

	// 支持 preload，点号表示关联，如 user.username
	if strings.Contains(field, ".") {
		parts := strings.Split(field, ".")
		assoc := parts[0]
		assocField := strings.Join(parts[1:], ".")
		// preload 关联
		db = db.Preload(assoc)
		// 这里假设关联表别名是复数，实际根据数据库设计调整
		alias := assoc + "s"

		// JSON字段（user.meta->>'key'）
		if strings.Contains(assocField, "->>") {
			jsonParts := strings.Split(assocField, "->>")
			baseField := jsonParts[0]
			jsonKey := strings.Trim(jsonParts[1], "'")

			switch op {
			case "=":
				db = db.Where(fmt.Sprintf("%s.%s->>'%s' = ?", alias, baseField, jsonKey), value)
			case "!=":
				db = db.Where(fmt.Sprintf("%s.%s->>'%s' <> ?", alias, baseField, jsonKey), value)
			default:
				db = db.Where(fmt.Sprintf("%s.%s->>'%s' %s ?", alias, baseField, jsonKey, op), value)
			}
			return db
		}

		// 普通字段
		switch op {
		case "=":
			db = db.Where(fmt.Sprintf("%s.%s = ?", alias, assocField), value)
		case "like":
			db = db.Where(fmt.Sprintf("%s.%s LIKE ?", alias, assocField), value)
		case "in":
			db = db.Where(fmt.Sprintf("%s.%s IN ?", alias, assocField), value)
		case "not in":
			db = db.Where(fmt.Sprintf("%s.%s NOT IN ?", alias, assocField), value)
		case "!=":
			db = db.Where(fmt.Sprintf("%s.%s <> ?", alias, assocField), value)
		default:
			db = db.Where(fmt.Sprintf("%s.%s %s ?", alias, assocField, op), value)
		}
		return db
	}

	// 处理 JSON 字段查询，例： meta->>'key'
	if strings.Contains(field, "->>") {
		parts := strings.Split(field, "->>")
		baseField := parts[0]
		jsonKey := strings.Trim(parts[1], "'")
		switch op {
		case "=":
			db = db.Where(fmt.Sprintf("%s->>'%s' = ?", baseField, jsonKey), value)
		case "!=":
			db = db.Where(fmt.Sprintf("%s->>'%s' <> ?", baseField, jsonKey), value)
		case "in":
			db = db.Where(fmt.Sprintf("%s->>'%s' IN ?", baseField, jsonKey), value)
		default:
			db = db.Where(fmt.Sprintf("%s->>'%s' %s ?", baseField, jsonKey, op), value)
		}
		return db
	}

	// 普通字段
	switch op {
	case "=":
		db = db.Where(fmt.Sprintf("%s = ?", field), value)
	case "!=":
		db = db.Where(fmt.Sprintf("%s <> ?", field), value)
	case ">":
		db = db.Where(fmt.Sprintf("%s > ?", field), value)
	case "<":
		db = db.Where(fmt.Sprintf("%s < ?", field), value)
	case ">=":
		db = db.Where(fmt.Sprintf("%s >= ?", field), value)
	case "<=":
		db = db.Where(fmt.Sprintf("%s <= ?", field), value)
	case "in":
		db = db.Where(fmt.Sprintf("%s IN ?", field), value)
	case "not in":
		db = db.Where(fmt.Sprintf("%s NOT IN ?", field), value)
	case "like":
		db = db.Where(fmt.Sprintf("%s LIKE ?", field), value)
	case "not like":
		db = db.Where(fmt.Sprintf("%s NOT LIKE ?", field), value)
	case "is null":
		db = db.Where(fmt.Sprintf("%s IS NULL", field))
	case "is not null":
		db = db.Where(fmt.Sprintf("%s IS NOT NULL", field))
	// ... 更多操作符
	default:
		db = db.Where(fmt.Sprintf("%s %s ?", field, op), value)
	}
	return db
}
