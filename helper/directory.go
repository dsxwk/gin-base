package helper

import (
	"errors"
	"gin-base/common/global"
	"go.uber.org/zap"
	"os"
)

// PathExists 文件目录是否存在
// @param: path string
// @return: bool, error
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
// @param: dirs ...string
// @return: err error
func CreateDir(dirs ...string) (err error) {
	var (
		exist bool
	)
	for _, v := range dirs {
		exist, err = PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.Log.Debug("create directory" + v)
			if err = os.MkdirAll(v, os.ModePerm); err != nil {
				global.Log.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
