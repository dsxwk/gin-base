package main

import (
	"fmt"
	"gin-base/common/global"
	"gin-base/config"
	"gin-base/helper/utils"
	"github.com/spf13/pflag"
	"gorm.io/gen"
	"gorm.io/gorm"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Options struct {
	Make        string // 文件类型
	Function    string // 函数方法
	Method      string // 请求方法
	Router      string // 路由
	Description string // 函数方法描述或注释
	TableName   string // 模型对应的表名
	Path        string // 模型生成路径
	Camel       bool   // 是否驼峰命名
	FileName    string // 生成的文件名
}

func main() {
	var (
		opts         Options
		templateFile string
		rootPath     = config.GetRootPath()
	)

	pflag.StringVarP(&opts.Make, "make", "m", "", "文件类型(示例：controller|model|service|validate|middleware|router)")
	pflag.StringVarP(&opts.Function, "function", "f", "", "函数方法(示例：create)")
	pflag.StringVarP(&opts.Method, "method", "e", "", "请求方法(示例：get)")
	pflag.StringVarP(&opts.Router, "router", "r", "", "路由(示例：/v1/user)")
	pflag.StringVarP(&opts.Description, "description", "d", "", "函数方法描述或注释")
	pflag.StringVarP(&opts.TableName, "tableName", "t", "", "模型对应的表名(示例：user)")
	pflag.StringVarP(&opts.Path, "path", "p", "", "模型生成路径(默认可不传)")
	pflag.BoolVarP(&opts.Camel, "camel", "c", false, "是否驼峰命名(示例：true 默认为false)")
	pflag.StringVarP(&opts.FileName, "fileName", "i", "", "生成的文件名(示例：/app/controller/v1/test)")

	pflag.Parse()

	if opts.Make == "" {
		fmt.Println("请输入正确的文件类型,Usage: go run ./cli/main.go --make=<controller|model|service|validate|middleware|router>")
		return
	}

	switch opts.Make {
	case "model":
		generateModel(opts.TableName, opts.Path, opts.Camel)
		break
	case "controller", "service", "validate", "middleware", "router":
		templateFile = filepath.Join(rootPath, "common", "template", opts.Make+".tpl")
		break
	default:
		fmt.Println("暂不支持make类型为【" + opts.Make + "】暂只支持命令类型为【controller, model, service, validate, middleware, router】")
		return
	}

	if opts.Make == "model" && opts.TableName == "" {
		fmt.Println("请输入模型对应的表名,Usage: go run ./cli/main.go --make=" + opts.Make + " --tableName=your table name")
		return
	}

	if opts.Make != "model" && opts.FileName == "" {
		fmt.Println("请输入文件名,Usage: go run ./cli/main.go --make=" + opts.Make + " --fileName=/app/your path/your file name")
		return
	}

	if err := generateFile(templateFile, opts.FileName, opts.Function, opts.Method, opts.Router, opts.Description); err != nil {
		fmt.Println(err.Error())
		return
	}
}

// generateFile 生成模版文件
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
		function = "FunctionName"
	}

	if method == "" {
		method = "get"
	}

	if description == "" {
		description = "Your Description"
	}

	if router == "" {
		router = "RouterName"
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

// generateModel 生成 model
func generateModel(tableName string, path string, camel bool) {
	if tableName == "" {
		fmt.Println("请输入表名,Usage: go run ./cli/main.go --make=model --tableName=your table name")
		return
	}

	if path == "" {
		path = filepath.Join(config.GetRootPath() + "/app/temp")
	}
	createTableStruct(tableName, path, camel)

	fmt.Println(tableName + " 表结构体生成成功!")
}

// createTableStruct 生成表结构体
// @param tableName 表名
// @param path 生成路径
func createTableStruct(tableName string, path string, camel bool) {
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
		"json":      func(detailType gorm.ColumnType) (dataType string) { return "JsonString" },
		"datetime": func(detailType gorm.ColumnType) (dataType string) {
			// 针对 deleted_at 字段特殊处理
			if detailType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}
			return "*JsonTime"
		},
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

	// 修改文件，插入钩子方法
	insertHooksIntoModel(path, tableName)

	_ = os.RemoveAll(path)
}

// insertHooksIntoModel 插入钩子方法到生成的模型文件
func insertHooksIntoModel(path string, tableName string) {
	// 确保路径和文件名正确
	modelFilePath := filepath.Join(config.GetRootPath(), "app", "model", tableName+".gen.go")
	structName := utils.ToCamelCase(tableName)
	// 首字母大写
	structName = strings.ToUpper(string(structName[0])) + structName[1:]

	// 读取文件内容
	content, err := ioutil.ReadFile(modelFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fileContent := string(content)

	// 检查是否已导入 time 包
	if !strings.Contains(fileContent, `"time"`) {
		// 用正则查找 import 块
		re := regexp.MustCompile(`import\s*\(([^)]*)\)`)
		if re.MatchString(fileContent) {
			// 已有多行 import，插入 "time"
			fileContent = re.ReplaceAllStringFunc(fileContent, func(s string) string {
				if strings.Contains(s, `"time"`) {
					return s
				}
				return strings.Replace(s, ")", `    "time"`+"\n)", 1)
			})
		} else {
			// 没有多行 import，查找单行 import 或插入新的 import 块
			if strings.Contains(fileContent, `import "gorm.io/gorm"`) {
				fileContent = strings.Replace(fileContent, `import "gorm.io/gorm"`, `import (
    "gorm.io/gorm"
    "time"
)`, 1)
			} else {
				// 没有 import，插入到 package 后
				lines := strings.SplitN(fileContent, "\n", 2)
				if len(lines) > 1 {
					fileContent = lines[0] + "\nimport (\n    \"time\"\n)\n" + lines[1]
				}
			}
		}
		// 覆盖写回文件
		err = ioutil.WriteFile(modelFilePath, []byte(fileContent), 0666)
		if err != nil {
			fmt.Printf("Error writing import to file: %v\n", err)
			return
		}
	}

	//	// 追加钩子方法
	//	file, err := os.OpenFile(modelFilePath, os.O_APPEND|os.O_RDWR, 0666)
	//	if err != nil {
	//		fmt.Printf("Error opening file: %v\n", err)
	//		return
	//	}
	//	defer func(file *os.File) {
	//		err = file.Close()
	//		if err != nil {
	//			fmt.Printf("Error closing file: %v\n", err)
	//			return
	//		}
	//	}(file)
	//
	//	// 准备钩子方法的内容
	//	hooks := fmt.Sprintf(`
	//// BeforeCreate 创建之前
	//func (s *%s) BeforeCreate(tx *gorm.DB) (err error) {
	//    now := JsonTime(time.Now())
	//    s.CreatedAt = &now
	//    s.UpdatedAt = &now
	//    return nil
	//}
	//
	//// BeforeUpdate 更新之前
	//func (s *%s) BeforeUpdate(tx *gorm.DB) (err error) {
	//    now := JsonTime(time.Now())
	//    s.UpdatedAt = &now
	//    return nil
	//}
	//`, structName, structName)
	//
	//	// 将钩子方法写入文件
	//	_, err = file.WriteString(hooks)
	//	if err != nil {
	//		fmt.Printf("Error writing hooks to file: %v\n", err)
	//		return
	//	}
	//
	//	fmt.Println("钩子方法已插入到模型文件:", modelFilePath)
}
