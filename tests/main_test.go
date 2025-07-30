package tests

import (
	"gin-base/config"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// 使用 runtime.Caller 获取当前文件路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("failed to get current file info")
	}

	projectRoot := filepath.Dir(filename)
	projectRoot = filepath.Dir(projectRoot)

	err := os.Chdir(projectRoot)
	if err != nil {
		log.Fatalf("failed to change dir to project root: %v", err)
	}

	log.Printf("Changed working directory to: %s", projectRoot)

	config.InitConfig()

	code := m.Run()
	os.Exit(code)
}
