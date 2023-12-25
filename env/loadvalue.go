package env

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func setUpZap() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/apilog.log",
		MaxSize:    30,  // megabytes
		MaxBackups: 300, // backup files
		MaxAge:     31,  // days
	})
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = "caller"
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.L().Info("zap started")

}
