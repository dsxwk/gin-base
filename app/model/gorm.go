package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type JsonTime time.Time

type JsonString []string

type JsonInt64 []int64

type BoolInt64 bool

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
