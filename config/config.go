package config

import (
	"errors"
	"fmt"
	"gin-base/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Log 日志
type Log struct {
	MaxSize    int `yaml:"max-size"`    // 单个日志文件大小（MB）
	MaxBackups int `yaml:"max-backups"` // 最多保留的旧日志文件数
	MaxDay     int `yaml:"max-day"`     // 保留的最大天数
}

// Jwt token
type Jwt struct {
	Key string `yaml:"key"`
	Exp int64  `yaml:"exp"`
}

// Cors 跨域
type Cors struct {
	AllowOrigin      string `yaml:"allow-origin"`
	AllowHeaders     string `yaml:"allow-headers"`
	ExposeHeaders    string `yaml:"expose-headers"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowCredentials string `yaml:"allow-credentials"`
}

// Cache 缓存
type Cache struct {
	Type  string `yaml:"type"`
	Redis Redis
}

// Config 配置
type Config struct {
	Mysql
	Log
	Jwt
	Cors
	Cache
}

var (
	config = InitConfig()
	logs   = InitLogger()
)

// InitConfig 初始化配置
func InitConfig() Config {
	rootPath := GetRootPath()
	file := filepath.Join(rootPath+"/config", "config.yaml")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	if err = yaml.Unmarshal(yamlFile, &c); err != nil {
		log.Fatal(err)
	}
	return c
}

// InitCache 初始化缓存
func InitCache() helper.CacheInterface {
	var (
		c helper.CacheInterface
	)

	// 没有配置缓存类型,默认使用内存缓存
	if config.Cache.Type == "memory" || config.Cache.Type == "" {
		c = helper.NewMemoryCache(5*time.Minute, 10*time.Minute)
	} else if config.Cache.Type == "redis" {
		// 使用 Redis 缓存
		c = InitRedis()
	} else {
		log.Fatalf("Unsupported cache type: %s", config.Cache.Type)
	}

	fmt.Println("Cache initialized with type:", config.Cache.Type)

	return c
}

// InitDB 初始化数据库
func InitDB() *gorm.DB {
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

// InitLogger 初始化日志
func InitLogger() *zap.Logger {
	// 设置日志文件输出路径和文件年月日
	logPath := "log/" + time.Now().Format("2006-01-02") + ".log"

	// 配置 Lumberjack 日志轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    config.MaxSize,    // 单个日志文件大小（MB）
		MaxBackups: config.MaxBackups, // 最多保留的旧日志文件数
		MaxAge:     config.MaxDay,     // 保留的最大天数
		Compress:   true,
	}

	// 创建一个Zap配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberJackLogger),
		atomicLevel,
	)

	return zap.New(core)
}

// InitRedis 初始化redis
func InitRedis() *helper.RedisCache {
	return helper.NewRedisCache(config.Cache.Redis.Address, config.Cache.Redis.Password, config.Cache.Redis.DB)
}

// GetRootPath 获取根目录
func GetRootPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	adsPath := strings.Replace(pwd, "\\", "/", -1)

	return adsPath
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
		// 获取查询开始时间
		startTime, _ := db.InstanceGet("gorm:query_slowquery_start_time")
		if startTime != nil {
			elapsed := time.Since(startTime.(time.Time))
			// 超过阈值则记录慢查询
			if elapsed > time.Millisecond*config.Mysql.SlowQuerySeconds {
				sql := db.Statement.SQL.String()
				// 记录慢查询的执行时间和查询语句
				logs.Warn("执行慢查询："+elapsed.String(), zap.Error(errors.New("执行慢查询："+elapsed.String()+" sql："+sql)))
			}
		}
	})
}
