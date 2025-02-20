package config

import (
	"fmt"
	"gin-base/helper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// mysql
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

// 日志
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

// 跨域
type Cors struct {
	AllowOrigin      string `yaml:"allow-origin"`
	AllowHeaders     string `yaml:"allow-headers"`
	ExposeHeaders    string `yaml:"expose-headers"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowCredentials string `yaml:"allow-credentials"`
}

// 缓存
type Cache struct {
	Type  string `yaml:"type"`
	Redis Redis
}

// 配置
type Config struct {
	Mysql
	Log
	Jwt
	Cors
	Cache
}

// 初始化配置
func InitConfig() Config { // 初始化数据
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

// 初始化缓存
func InitCache() helper.CacheInterface {
	var (
		c      helper.CacheInterface
		config = InitConfig()
	)

	// 没有配置缓存类型,默认使用内存缓存
	if config.Cache.Type == "memory" || config.Cache.Type == "" {
		c = helper.NewMemoryCache(5*time.Minute, 10*time.Minute)
	} else if config.Cache.Type == "redis" {
		// 使用 Redis 缓存
		c = helper.NewRedisCache(config.Cache.Redis.Address, config.Cache.Redis.Password, config.Cache.Redis.DB)
	} else {
		log.Fatalf("Unsupported cache type: %s", config.Cache.Type)
	}

	fmt.Println("Cache initialized with type:", config.Cache.Type)

	return c
}

// 获取根目录
func GetRootPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	adsPath := strings.Replace(pwd, "\\", "/", -1)

	return adsPath
}
