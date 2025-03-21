package config

import (
	"errors"
	"fmt"
	"gin-base/common/extend/event"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	SQLRes []string
	mu     sync.Mutex // 用于保护 SQLRes 的并发安全
)

// Mysql mysql
type Mysql struct {
	Host             string        `yaml:"host"`
	Port             string        `yaml:"port"`
	Name             string        `yaml:"name"`
	User             string        `yaml:"user"`
	Password         string        `yaml:"password"`
	SlowQuerySeconds time.Duration `yaml:"slow-query-seconds"`
}

// InitMysql 初始化数据库
func InitMysql() *gorm.DB {
	var (
		err error
		DB  *gorm.DB
	)
	DB, err = gorm.Open(mysql.Open(config.Mysql.Host), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 设置log输出到控制台
			logger.Config{
				SlowThreshold: time.Millisecond * config.Mysql.SlowQuerySeconds, //慢SQL阈值（单位：ms）
				LogLevel:      logger.Info,
				// IgnoreRecordNotFoundError: true, // 忽略Record Not Found 错误
				Colorful: true, // 给log添加颜色
			},
		),
	})
	if err != nil {
		logs.Error(err.Error(), zap.Error(err))
		panic("数据库链接错误")
	}
	q, err := DB.DB()
	q.SetMaxIdleConns(10)
	q.SetMaxOpenConns(100)
	q.SetConnMaxLifetime(59 * time.Second)

	// 日志记录慢查询
	LogSlowQuery(DB)

	return DB
}

// AddSQL 线程安全地追加 SQL 数据
func AddSQL(sql string, vars []interface{}) {
	mu.Lock()
	defer mu.Unlock()

	SQLRes = append(SQLRes, getSQL(sql, vars))
}

// GetSQLLogs 获取所有 SQL 日志记录
func GetSQLLogs() []string {
	mu.Lock()
	defer mu.Unlock()

	// 复制一份 SQLRes 返回，避免外部修改
	sqlLogs := make([]string, len(SQLRes))
	copy(sqlLogs, SQLRes)

	return sqlLogs
}

// LogSlowQuery 日志记录慢查询
func LogSlowQuery(DB *gorm.DB) {
	// 注册查询前回调
	_ = DB.Callback().Query().Before("gorm:query").Register("slowquery:begin", func(db *gorm.DB) {
		// 记录查询开始时间
		db.InstanceSet("gorm:query_slowquery_start_time", time.Now())
	})

	// 注册查询后回调
	_ = DB.Callback().Query().After("gorm:query").Register("slowquery:end", func(db *gorm.DB) {
		var (
			sql  = db.Statement.SQL.String()
			vars = db.Statement.Vars
		)

		// 追加 SQL 数据
		AddSQL(sql, vars)

		// 发布 SQL 事件
		sqlEvent := event.Event{
			Name: "sql_listen",
			Data: SQLRes,
		}
		EventBus.Publish(sqlEvent)

		// 获取查询开始时间
		startTime, _ := db.InstanceGet("gorm:query_slowquery_start_time")
		if startTime != nil {
			elapsed := time.Since(startTime.(time.Time))
			// 超过阈值则记录慢查询
			if elapsed > time.Millisecond*config.Mysql.SlowQuerySeconds {
				// 记录慢查询的执行时间和查询语句
				logs.Warn("执行慢查询："+elapsed.String(), zap.Error(errors.New("执行慢查询："+elapsed.String()+" sql："+sql)))
			}
		}
	})
}

// getSQL 替换 SQL 中的占位符 "?" 为实际值
func getSQL(sql string, vars []interface{}) string {
	for _, v := range vars {
		// 将参数值格式化为字符串
		var formattedValue string
		switch value := v.(type) {
		case string:
			// 如果是字符串，加上引号
			formattedValue = fmt.Sprintf("'%s'", value)
		case []byte:
			// 如果是字节数组，转换为字符串并加上引号
			formattedValue = fmt.Sprintf("'%s'", string(value))
		default:
			// 其他类型直接转换为字符串
			formattedValue = fmt.Sprintf("%v", value)
		}

		// 替换第一个 "?" 为实际值
		sql = strings.Replace(sql, "?", formattedValue, 1)
	}
	return sql
}
