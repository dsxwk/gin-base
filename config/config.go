package config

import (
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

// 日志
type Log struct {
	MaxSize    int `yaml:"max-size"`    // 单个日志文件大小（MB）
	MaxBackups int `yaml:"max-backups"` // 最多保留的旧日志文件数
	MaxDay     int `yaml:"max-day"`     // 保留的最大天数
}

// 配置
type config struct {
	Mysql
	Log
}

// 初始化配置
func InitConfig() config { // 初始化数据
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	adsPath := strings.Replace(pwd, "\\", "/", -1)
	file := filepath.Join(adsPath+"/config", "config.yaml")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	c := config{}
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		log.Fatal(err)
	}
	return c
}