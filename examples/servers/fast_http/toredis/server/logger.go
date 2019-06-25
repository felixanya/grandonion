package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var logger *zap.Logger

type Logger struct {
	LogPath 		string
	MaxSize 		int
	MaxBackups 		int
	MaxAge 			int
	Level			string
}

func (lc *Logger) initLogger() error {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: lc.LogPath,
		MaxSize: lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge: lc.MaxAge,
	})

	encoderConf := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level := zap.DebugLevel
	levelLower := strings.ToLower(lc.Level)
	switch levelLower {
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf), w, level)
	logger = zap.New(core)
	return nil
}

