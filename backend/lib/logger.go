// TODO: add file rotate using  github.com/lestrrat-go/file-rotatelogs
package lib

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

func GetEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  singleton.Config.Zap.StacktraceKey,
		EncodeTime:     GetCustomTimeEncoder, // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   GetCustomCallerEncoder, // caller trace encoder
	}
	GetEncodeLevel(&config)
	return config
}

// 自定义文件：行号输出项
func GetCustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

// Set logger encode level based on config
func GetEncodeLevel(c *zapcore.EncoderConfig) {
	switch singleton.Config.Zap.EncodeLevel {
	case "LowercaseLevelEncoder":
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		c.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		c.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
}

// Set CustomTimeEncoder
func GetCustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(singleton.Config.Zap.Prefix + " 2006/01/02 - 15:04:05.000 "))
}

func GetOutputPath() string {
	if ok, _ := PathExists(singleton.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", singleton.Config.Zap.Director)
		_ = os.Mkdir(singleton.Config.Zap.Director, os.ModePerm)
	}
	return fmt.Sprintf("./%s/%s", singleton.Config.Zap.Director, singleton.Config.Zap.FileName)
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("file already exist")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func LoggerInit() *zap.Logger {
	encoderConfig := GetEncoderConfig()

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Level:       atom,  // log level
		Development: false, // 开发模式，堆栈跟踪

		// DisableStacktrace completely disables automatic stacktrace capturing. By
		// default, stacktraces are captured for WarnLevel and above logs in
		// development and ErrorLevel and above in production.
		DisableStacktrace: true,
		Encoding:          singleton.Config.Zap.Format, // console or json
		EncoderConfig:     encoderConfig,
		InitialFields:     map[string]interface{}{"serviceName": "wsl"}, // add new key-value pair
		OutputPaths:       []string{"stdout", GetOutputPath()},          //  stdout and customized destination
		ErrorOutputPaths:  []string{"stderr"},
	}

	// build log
	ZapLog_V1, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed log initialization : %v", err))
	}

	// testing
	ZapLog_V1.Warn("warn: log initialization successful")
	ZapLog_V1.Info("info: log initialization successful")
	ZapLog_V1.Error("err: log initialization successful")

	// print out error
	// if _, err := strconv.Atoi("a123"); err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	ZapLog_V1.Error(fmt.Sprintf("err: %v", err))
	// }
	return ZapLog_V1
}
