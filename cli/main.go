package main

import (
	"flag"
	"fmt"
	"gin-base/common/global"
	"gin-base/config"
	"gin-base/helper/utils"
	"gorm.io/gen"
	"gorm.io/gorm"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		templateFile string
		rootPath     = config.GetRootPath()
	)

	// 文件类型
	_make := flag.String("make", "", "文件类型")

	// 函数
	function := flag.String("function", "", "函数")

	// 方法
	method := flag.String("method", "", "方法")

	// 路由
	router := flag.String("router", "", "路由")

	// 描述
	description := flag.String("description", "", "描述")

	// 表名,默认为空,描述
	tableName := flag.String("tableName", "", "表名")

	// 生成路径,默认为空,描述
	path := flag.String("path", "", "生成路径")

	// 是否生成驼峰字段
	camel := flag.Bool("camel", false, "是否生成驼峰字段")

	// 生成的文件名
	fileName := flag.String("fileName", "", "文件名")

	// 解析命令行参数
	flag.Parse()

	if *_make == "" {
		fmt.Println("请输入正确的文件类型,Usage: go run main.go --make=<controller|model|service|validate|middleware>")
		return
	}

	switch *_make {
	case "model":
		generateModel(*tableName, *path, *camel)
		break
	case "controller":
		templateFile = filepath.Join(rootPath, "common", "template", *_make+".tpl")
		break
	case "service":
		templateFile = filepath.Join(rootPath, "common", "template", *_make+".tpl")
		break
	case "validate":
		templateFile = filepath.Join(rootPath, "common", "template", *_make+".tpl")
		break
	case "middleware":
		templateFile = filepath.Join(rootPath, "common", "template", *_make+".tpl")
		break
	default:
		fmt.Println("暂不支持make类型为【" + *_make + "】暂只支持命令类型为【controller, model, service, validate, middleware】")
		break
	}

	if *fileName == "" {
		fmt.Println("请输入文件名,Usage: go run main.go --make=" + *_make + " --fileName=/app/your path/your file name")
		return
	}

	if err := generateFile(templateFile, *fileName, *function, *method, *router, *description); err != nil {
		fmt.Println(err.Error())
		return
	}
}

// 生成模版文件
func generateFile(templateFile, filename string, function string, method string, router string, description string) error {
	// 加载模板文件
	fmt.Printf("Loading template file: %s\n", templateFile)
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error parsing template:", err.Error())
		return err
	}

	// 提取包名（即文件路径中的最后一个目录作为包名）
	packageName := filepath.Base(filepath.Dir(filename))
	fmt.Printf("Detected package name: %s\n", packageName)

	// 设置默认值
	if function == "" {
		function = "YourFunctionName"
	}

	if method == "" {
		method = "get|post ..."
	}

	if description == "" {
		description = "YourDescription"
	}

	if router == "" {
		router = "YourRouter"
	}

	// 确保目录存在
	dir := filepath.Dir(filename)
	fmt.Printf("Checking if directory exists: %s\n", dir)

	// 使用 os.Stat 检查目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist. Creating: %s\n", dir)
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Failed to create directory:", err.Error())
			return err
		}
	} else {
		fmt.Println("Directory already exists.")
	}

	// 创建文件
	filePath := filepath.Join(config.GetRootPath(), filename+".go")
	fmt.Printf("Creating file: %s\n", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file:", err.Error())
		return err
	}
	defer file.Close()

	// 准备数据
	data := struct {
		Package     string
		Name        string
		Function    string
		Router      string
		Method      string
		Description string
	}{
		Package:     packageName,                            // 提取的包名
		Name:        strings.Title(filepath.Base(filename)), // 控制器名称（首字母大写）
		Function:    function,                               // 如果为空，使用默认值
		Router:      router,                                 // 如果为空，使用默认值
		Method:      method,                                 // 如果为空，使用默认值
		Description: description,                            // 如果为空，使用默认值
	}

	// 执行模板并写入文件
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		return err
	} else {
		fmt.Println("Template executed and content written to file.")
	}

	fmt.Println(filePath + " 生成成功!")
	return nil
}

// 生成 model
func generateModel(tableName string, path string, camel bool) {
	if tableName == "" {
		fmt.Println("请输入表名,Usage: go run main.go --make=model --tableName=your table name")
		return
	}

	if path == "" {
		path = "app/temp"
	}
	createTableStruct(tableName, path, camel)

	fmt.Println(tableName + " 表结构体生成成功!")
}

// 生成表结构体
// @param tableName 表名
// @param path 生成路径
func createTableStruct(tableName string, path string, camel bool) {
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
		"datetime":  func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 datetime 转换为 string
		//"timestamp":  func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 timestamp 转换为 string
		//"date":       func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 date 转换为 string
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义JSON tag
	if camel {
		g.WithJSONTagNameStrategy(func(columnName string) string {
			return utils.ToCamelCase(columnName)
		})
	}

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	TargetTable := g.GenerateModel(tableName)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(TargetTable)

	g.Execute()

	_ = os.RemoveAll(path)
}
