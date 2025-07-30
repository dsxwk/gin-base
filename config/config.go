package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// Config 配置
type Config struct {
	Mysql
	Log
	Jwt
	Cors
	Cache
	Service
	Event
}

// InitConfig 初始化配置
func InitConfig() *Config {
	file := filepath.Join(GetRootPath()+"/config", "config.yaml")
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	c := &Config{}
	if err = yaml.Unmarshal(yamlFile, &c); err != nil {
		log.Fatal(err)
	}
	return c
}

// GetRootPath 获取根目录
//func GetRootPath() string {
//	pwd, err := os.Getwd()
//	if err != nil {
//		panic(err)
//		os.Exit(1)
//	}
//	adsPath := strings.Replace(pwd, "\\", "/", -1)
//
//	return adsPath
//}

// GetRootPath 获取项目根路径
func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	rootPath := filepath.Join(basePath, "..")
	rootPath, _ = filepath.Abs(rootPath)

	return strings.ReplaceAll(rootPath, "\\", "/")
}
