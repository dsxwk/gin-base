package main

import (
	"fmt"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/config"
	"gin-base/helper/utils"
	"github.com/spf13/pflag"
	"gorm.io/gen"
	"gorm.io/gorm"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type CliCommand struct {
	common.BaseCommand
}

type Options struct {
	Make        string // 文件类型 controller|model|router ...
	Function    string // 函数方法 List|Create ...
	Method      string // 请求方法 get|post ...
	Router      string // 路由 /v1/user ...
	Description string // 函数方法描述或注释 列表|创建 ...
	Table       string // 模型对应的表名 user|article ...
	File        string // 路径文件 /v1/user ... 注：将在对应的类型目录下生成 v1/user.go 文件,如为controller 则生成 root/app/controller/v1/user.go 文件
	Camel       bool   // 是否驼峰命名
}

var (
	rootPath = config.GetRootPath()
)

func main() {
	var (
		s            CliCommand
		opts         Options
		templateFile string
	)

	// 检查命令是否提供
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cli/main.go <command> [--file=<path>]")
		return
	}

	// 获取主命令（如 make:controller 或 make:model）
	command := os.Args[1]
	if !strings.HasPrefix(command, "make:") {
		fmt.Println("Error: Command must start with 'make:'")
		return
	}

	opts.Make = strings.TrimPrefix(command, "make:") // 文件类型(示例：controller|model|service|validate|middleware|router)

	pflag.StringVarP(&opts.Function, "func", "c", "", "函数方法(示例：create)")
	pflag.StringVarP(&opts.Method, "method", "m", "", "请求方法(示例：get)")
	pflag.StringVarP(&opts.Router, "router", "r", "", "路由(示例：/v1/user)")
	pflag.StringVarP(&opts.Description, "desc", "d", "", "函数方法描述或注释")
	pflag.StringVarP(&opts.Table, "table", "t", "", "模型对应的表名(示例：user)")
	pflag.StringVarP(&opts.File, "file", "f", "", "生成的路径文件(示例：v1/test)")
	pflag.BoolVarP(&opts.Camel, "camel", "l", false, "是否驼峰命名(示例：true 默认为false)")

	pflag.Parse()

	if opts.Make == "" {
		fmt.Println("请输入正确的命令,Usage: go run ./cli/main.go make:controller|model|service|validate|middleware|router")
		return
	}

	switch opts.Make {
	case "model":
		s.generateModel(opts.Table, opts.File, opts.Camel)
		break
	case "controller", "service", "validate", "middleware", "router":
		templateFile = filepath.Join(rootPath, "common", "template", opts.Make+".tpl")
		break
	default:
		fmt.Println("暂不支持make类型为【" + opts.Make + "】暂只支持命令类型为【controller, model, service, validate, middleware, router】")
		return
	}

	if opts.Make == "model" && opts.Table == "" {
		fmt.Println("请输入模型对应的表名,Usage: go run ./cli/main.go make:" + opts.Make + " --table=YourTable")
		return
	}

	if opts.Make != "model" && opts.File == "" {
		fmt.Println("请输入文件名,Usage: go run ./cli/main.go make:" + opts.Make + " --file=YourPath")
		return
	}

	if opts.File != "" {
		opts.File = s.getMakePath(opts.File, opts.Make)
	}

	if err := s.generateFile(templateFile, opts); err != nil {
		fmt.Println(err.Error())
		return
	}
}

// getMakePath 获取make文件路径
func (s CliCommand) getMakePath(file string, cType string) string {
	if !strings.HasSuffix(file, "/") {
		file = "/" + file
	}

	switch cType {
	case "router":
		file = filepath.Join("/app", "routers", file)
	default:
		file = filepath.Join("/app", cType, file)
	}

	return file
}

// generateFile 生成模版文件
func (s CliCommand) generateFile(templateFile string, opts Options) error {
	// 加载模板文件
	fmt.Printf("Loading template file: %s\n", templateFile)
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error parsing template:", err.Error())
		return err
	}

	// 提取包名（即文件路径中的最后一个目录作为包名）
	packageName := filepath.Base(filepath.Dir(opts.File))
	fmt.Printf("Detected package name: %s\n", packageName)

	// 设置默认值
	if opts.Function == "" {
		opts.Function = "YourFuncName"
	}

	if opts.Method == "" {
		opts.Method = "get"
	}

	if opts.Description == "" {
		opts.Description = "YourDesc"
	}

	if opts.Router == "" {
		opts.Router = "YourRouterName"
	}

	// 确保目录存在
	dir := filepath.Dir(rootPath + opts.File)
	fmt.Printf("Checking if directory exists: %s\n", dir)

	// 使用 os.Stat 检查目录是否存在
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist. Creating: %s\n", dir)
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Failed to create directory:", err.Error())
			return err
		}
	} else {
		fmt.Println("Directory already exists.")
	}

	// 创建文件
	file := filepath.Join(rootPath, opts.File+".go")
	fmt.Printf("Creating file: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Failed to create file:", err.Error())
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("Failed to close file:", err.Error())
		}
	}(f)

	// 准备数据
	data := struct {
		Package     string
		Name        string
		Function    string
		Router      string
		Method      string
		Description string
	}{
		Package:     packageName,                             // 提取的包名
		Name:        strings.Title(filepath.Base(opts.File)), // 控制器名称（首字母大写）
		Function:    opts.Function,                           // 如果为空，使用默认值
		Router:      opts.Router,                             // 如果为空，使用默认值
		Method:      opts.Method,                             // 如果为空，使用默认值
		Description: opts.Description,                        // 如果为空，使用默认值
	}

	// 执行模板并写入文件
	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		return err
	} else {
		fmt.Println("Template executed and content written to file.")
	}

	fmt.Println(file + " 生成成功!")
	return nil
}

// generateModel 生成 model
func (s CliCommand) generateModel(table string, file string, camel bool) {
	if table == "" {
		fmt.Println("请输入表名,Usage: go run ./cli/main.go make:model --table=YourTable")
		return
	}

	s.createTableStruct(table, file, camel)

	fmt.Println(table + " 表结构体生成成功!")
}

// createTableStruct 生成表结构体
// @param table 表名
// @param path 生成路径
func (s CliCommand) createTableStruct(table string, file string, camel bool) {
	var (
		outPath  = filepath.Join(rootPath + "/app/temp")
		fileName string
	)
	if file != "" {
		lastSlash := strings.LastIndex(file, "/")
		if lastSlash == -1 {
			fileName = file
			file = ""
		} else {
			fileName = file[lastSlash+1:]
			file = file[:lastSlash]
		}
	}

	file = filepath.Join(rootPath, "app", "model", file)

	g := gen.NewGenerator(gen.Config{
		OutPath:           outPath,
		Mode:              gen.WithDefaultQuery,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
		ModelPkgPath:      file,
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
		// "timestamp":  func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 timestamp 转换为 string
		// "date":       func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 date 转换为 string
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
	TargetTable := g.GenerateModel(table)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(TargetTable)

	g.Execute()

	// 修改文件，插入钩子方法
	// s.insertHooksIntoModel(table)

	// 修改文件名
	if file != filepath.Join(rootPath, "app", "model") || fileName != "" {
		s.modelHooks(table, file, fileName)
	}

	_ = os.RemoveAll(outPath)
}

// modelHooks 修改文件名以及更新模型
func (s CliCommand) modelHooks(table string, file string, fileName string) {
	// 修改文件名
	oldFilePath := filepath.Join(file, table+".gen.go")
	newFilePath := filepath.Join(file, fileName+".gen.go")
	// 检查旧文件是否存在
	if _, err := os.Stat(oldFilePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist, skipping rename.\n", oldFilePath)
	} else {
		// 重命名文件
		err := os.Rename(oldFilePath, newFilePath)
		if err != nil {
			fmt.Printf("Failed to rename file: %v\n", err)
		} else {
			fmt.Printf("Renamed %s to %s\n", oldFilePath, newFilePath)
		}
	}

	// 检查 gorm.go 文件是否存在
	gormFilePath := filepath.Join(file, "gorm.go")
	if _, err := os.Stat(gormFilePath); os.IsNotExist(err) {
		var (
			packageName string
		)
		// 将路径中的 '\' 替换为 '/'
		file = strings.ReplaceAll(file, "\\", "/")
		lastSlash := strings.LastIndex(file, "/")
		if lastSlash == -1 {
			packageName = file
		} else {
			packageName = file[lastSlash+1:]
		}

		// 创建 gorm.go 文件并写入内容
		content := `package ` + packageName + `

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type JsonTime time.Time

type JsonString []string

type JsonInt64 []int64

type BoolInt64 bool

// MarshalJSON 模型时间格式化公共方法
func (t *JsonTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(` + `""` + `), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JsonTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*t = JsonTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = JsonTime(v)
		return nil
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
		return nil
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
		return nil
	default:
		return fmt.Errorf("cannot convert %v to timestamp", value)
	}
}

func (j JsonString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonString) Scan(value interface{}) error {
	if value == nil {
		*j = JsonString{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonString", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j JsonInt64) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonInt64) Scan(value interface{}) error {
	if value == nil {
		*j = JsonInt64{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonInt64", value)
	}
	return json.Unmarshal(bytes, j)
}

// Value 写入数据库时：true → 1，false → 2
func (b BoolInt64) Value() (driver.Value, error) {
	if b {
		return 1, nil
	}
	return 2, nil
}

// Scan 读取数据库时：1 → true，2 → false（默认 false）
func (b *BoolInt64) Scan(value interface{}) error {
	if value == nil {
		*b = false
		return nil
	}
	var v int64
	switch val := value.(type) {
	case int64:
		v = val
	case int:
		v = int64(val)
	case int32:
		v = int64(val)
	case uint8:
		v = int64(val)
	case []byte:
		i, err := strconv.Atoi(string(val))
		if err != nil {
			return err
		}
		v = int64(i)
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		v = int64(i)
	default:
		return fmt.Errorf("unsupported Scan type for BoolInt12: %T", value)
	}

	*b = v == 1
	return nil
}
`
		err = os.WriteFile(gormFilePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", gormFilePath, err)
		} else {
			fmt.Printf("Created file %s with default content.\n", gormFilePath)
		}
	} else {
		fmt.Printf("File %s already exists, skipping creation.\n", gormFilePath)
	}
}

//// insertHooksIntoModel 插入钩子方法到生成的模型文件
//func (s CliCommand) insertHooksIntoModel(table string) {
//	// 确保路径和文件名正确
//	modelFilePath := filepath.Join(rootPath, "app", "model", table+".gen.go")
//	structName := utils.ToCamelCase(table)
//	// 首字母大写
//	structName = strings.ToUpper(string(structName[0])) + structName[1:]
//
//	// 读取文件内容
//	content, err := ioutil.ReadFile(modelFilePath)
//	if err != nil {
//		fmt.Printf("Error reading file: %v\n", err)
//		return
//	}
//	fileContent := string(content)
//
//	// 检查是否已导入 time 包
//	if !strings.Contains(fileContent, `"time"`) {
//		// 用正则查找 import 块
//		re := regexp.MustCompile(`import\s*\(([^)]*)\)`)
//		if re.MatchString(fileContent) {
//			// 已有多行 import，插入 "time"
//			fileContent = re.ReplaceAllStringFunc(fileContent, func(s string) string {
//				if strings.Contains(s, `"time"`) {
//					return s
//				}
//				return strings.Replace(s, ")", `    "time"`+"\n)", 1)
//			})
//		} else {
//			// 没有多行 import，查找单行 import 或插入新的 import 块
//			if strings.Contains(fileContent, `import "gorm.io/gorm"`) {
//				fileContent = strings.Replace(fileContent, `import "gorm.io/gorm"`, `import (
//   "gorm.io/gorm"
//   "time"
//)`, 1)
//			} else {
//				// 没有 import，插入到 package 后
//				lines := strings.SplitN(fileContent, "\n", 2)
//				if len(lines) > 1 {
//					fileContent = lines[0] + "\nimport (\n    \"time\"\n)\n" + lines[1]
//				}
//			}
//		}
//		// 覆盖写回文件
//		err = ioutil.WriteFile(modelFilePath, []byte(fileContent), 0666)
//		if err != nil {
//			fmt.Printf("Error writing import to file: %v\n", err)
//			return
//		}
//	}
//
//		// 追加钩子方法
//		file, err := os.OpenFile(modelFilePath, os.O_APPEND|os.O_RDWR, 0666)
//		if err != nil {
//			fmt.Printf("Error opening file: %v\n", err)
//			return
//		}
//		defer func(file *os.File) {
//			err = file.Close()
//			if err != nil {
//				fmt.Printf("Error closing file: %v\n", err)
//				return
//			}
//		}(file)
//
//		// 准备钩子方法的内容
//		hooks := fmt.Sprintf(`
//	// BeforeCreate 创建之前
//	func (s *%s) BeforeCreate(tx *gorm.DB) (err error) {
//	   now := JsonTime(time.Now())
//	   s.CreatedAt = &now
//	   s.UpdatedAt = &now
//	   return nil
//	}
//
//	// BeforeUpdate 更新之前
//	func (s *%s) BeforeUpdate(tx *gorm.DB) (err error) {
//	   now := JsonTime(time.Now())
//	   s.UpdatedAt = &now
//	   return nil
//	}
//	`, structName, structName)
//
//		// 将钩子方法写入文件
//		_, err = file.WriteString(hooks)
//		if err != nil {
//			fmt.Printf("Error writing hooks to file: %v\n", err)
//			return
//		}
//
//		fmt.Println("钩子方法已插入到模型文件:", modelFilePath)
//}
