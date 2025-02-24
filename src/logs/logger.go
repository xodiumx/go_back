package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Logger Global
var Logger *zap.SugaredLogger

// InitLogger initialize global logger
func InitLogger() {

	// Castom format for logs
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",                 // Rename ts → timestamp
		LevelKey:      "severity",                  // level → severity
		CallerKey:     "caller",                    // Which func is called current log caller'а
		MessageKey:    "message",                   // msg → message
		StacktraceKey: "stacktrace",                // Field for stacktrace
		EncodeTime:    zapcore.ISO8601TimeEncoder,  // Time format ISO8601 (2024-02-24T15:04:05Z)
		EncodeLevel:   zapcore.CapitalLevelEncoder, // INFO instead info
		EncodeCaller:  zapcore.ShortCallerEncoder,  // Short path (main.go:10)
	}
	// Addd color for log
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Change time format
	//encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	// Create custom encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel), // In stdout
	)

	// Create logger
	// AddCallerSkip - skip one n levels of func caller
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	Logger = logger.Sugar()
}

// CloseLogger clear resources (called after app close)
func CloseLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
