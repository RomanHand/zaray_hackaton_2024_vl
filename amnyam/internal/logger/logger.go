package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(levelStr string, debugMode bool) (*zap.Logger, error) {
	if err := validateLogLevel(levelStr); err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level := zap.NewAtomicLevel()
	level.SetLevel(convertStringLevelToZap(levelStr))

	config := zap.Config{
		Level:       level,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json", //json,console
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if debugMode {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
		config.Encoding = "console"
		config.Development = true
		config.Sampling = nil
	}

	//var err error
	lgr, err := config.Build()
	if err != nil {
		return nil, err
	}

	return lgr, nil
}

func validateLogLevel(level string) error {
	for _, l := range []string{"ERROR", "WARN", "INFO", "DEBUG"} {
		if l == level {
			return nil
		}
	}

	return fmt.Errorf("logger: wrong log level")
}

func convertStringLevelToZap(level string) zapcore.Level {
	switch level {
	case "ERROR":
		return zap.ErrorLevel
	case "WARN":
		return zap.WarnLevel
	case "INFO":
		return zap.InfoLevel
	case "DEBUG":
		return zap.DebugLevel
	default:
		return zap.ErrorLevel
	}
}
