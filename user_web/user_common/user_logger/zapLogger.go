package user_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"

	"py_gomall/v2/user_web/user_config"
)

func NewLogger() (logger *zap.Logger) {
	if user_config.AppConf.LogConf == nil {
		log.Printf("WARNING:: No log parameters configured, default configuration enabled!")
		logConfig := &user_config.LogConf{
			Level:        "debug",
			Path:         "logs/user_web.log",
			MaxSizeMB:    20,
			MaxAgeDay:    30,
			MaxBackupDay: 7,
			Compress:     false,
		}
		user_config.AppConf.LogConf = logConfig
	}
	encoder := logEncoder()
	writer := logWriter()
	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(user_config.AppConf.LogConf.Level)); err != nil {
		log.Fatalf("Failed to initialize Zap log,error:%v\n", err)
	}
	var core zapcore.Core
	if user_config.AppConf.LogConf.Level == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writer, level)
	}
	caller := zap.AddCaller()
	lg := zap.New(core, caller)
	// 替换zap库中的全局Logger
	zap.ReplaceGlobals(lg)
	return lg
}

func logEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func logWriter() zapcore.WriteSyncer {
	lj := &lumberjack.Logger{
		Filename:   user_config.AppConf.LogConf.Path,
		MaxSize:    user_config.AppConf.LogConf.MaxSizeMB,
		MaxBackups: user_config.AppConf.LogConf.MaxBackupDay,
		MaxAge:     user_config.AppConf.LogConf.MaxAgeDay,
		Compress:   user_config.AppConf.LogConf.Compress,
	}
	return zapcore.AddSync(lj)
}
