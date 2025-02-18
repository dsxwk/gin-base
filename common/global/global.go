package global

import (
	"errors"
	"gin-base/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB     *gorm.DB // 数据库全局变量
	Log    *zap.Logger
	Config = config.InitConfig()
)

// 公共响应
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 公共分页数据返回
type PageData struct {
	//  总条数
	Total int64 `json:"total"`
	// 当前页
	Page int `json:"page"`
	// 每页条数
	PageSize int `json:"pageSize"`
	// 数据列表
	List interface{} `json:"list"`
}

// 错误码
const (
	// 成功
	Success = 0
	// 参数错误
	ArgsError = 400
	// 请求未授权
	Unauthorized = 401
	// 系统错误
	SystemError = 500
)

func init() {
	InitLogger()

	var err error
	DB, err = gorm.Open(mysql.Open(Config.Mysql.Host), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 设置log输出到控制台
			logger.Config{
				SlowThreshold: time.Millisecond * Config.Mysql.SlowQuerySeconds, //慢SQL阈值（单位：ms）
				LogLevel:      logger.Info,
				// IgnoreRecordNotFoundError: true, // 忽略Record Not Found 错误
				Colorful: true, // 给log添加颜色
			},
		),
	})
	if err != nil {
		Log.Error(err.Error(), zap.Error(err))
		panic("数据库链接错误")
	}
	q, err := DB.DB()
	q.SetMaxIdleConns(10)
	q.SetMaxOpenConns(100)
	q.SetConnMaxLifetime(59 * time.Second)

	// 日志记录慢查询
	LogSlowQuery(DB)
}

// 初始化日志
func InitLogger() {
	// 设置日志文件输出路径和文件年月日
	logPath := "log/" + time.Now().Format("2006-01-02") + ".log"

	// 配置 Lumberjack 日志轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    Config.MaxSize,    // 单个日志文件大小（MB）
		MaxBackups: Config.MaxBackups, // 最多保留的旧日志文件数
		MaxAge:     Config.MaxDay,     // 保留的最大天数
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

	Log = zap.New(core)
}

// 日志记录慢查询
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
			if elapsed > time.Millisecond*Config.Mysql.SlowQuerySeconds {
				sql := db.Statement.SQL.String()
				// 记录慢查询的执行时间和查询语句
				Log.Warn("执行慢查询："+elapsed.String(), zap.Error(errors.New("执行慢查询："+elapsed.String()+" sql："+sql)))
			}
		}
	})
}
