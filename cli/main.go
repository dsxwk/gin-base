package main

import (
	"flag"
	"fmt"
	"gin-base/common/global"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

func main() {
	// 定义几个变量，用于接收命令行的参数值
	var (
		// 表名
		tableName string
		// 生成路径
		path string
	)

	// 表名,默认为空,描述
	flag.StringVar(&tableName, "tableName", "", "表名")

	// 生成路径,默认为空,描述
	flag.StringVar(&path, "path", "", "生成路径")

	// 解析命令行参数
	flag.Parse()

	if tableName == "" {
		fmt.Println("表名tableName不能为空")
		return
	}

	CreateTableStruct(tableName, path)

	// 检查目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		global.Log.Error("目录"+path+"不存在", zap.Error(err))
		return
	}

	// 删除临时目录和文件
	err := os.RemoveAll(path)
	if err != nil {
		global.Log.Error("删除临时目录失败:", zap.Error(err))
		return
	}

	fmt.Println(path + " 临时目录删除成功!")
	fmt.Println(tableName + " 表结构体生成成功!")
}

// 生成表结构体
// @param tableName 表名
// @param path 生成路径
func CreateTableStruct(tableName string, path string) {
	if path == "" {
		path = "./query"
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           path,
		Mode:              gen.WithDefaultQuery,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(global.DB)
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	TargetTable := g.GenerateModel(tableName)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(TargetTable)

	g.Execute()
}
