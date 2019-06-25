package up2redis


import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var (
	logger *zap.Logger
	redisClient *redis.ClusterClient
)

type LoggerC struct {
	LogPath 		string
	MaxSize 		int
	MaxBackups 		int
	MaxAge 			int
	Level			string
}

type RedisC struct {
	DB				int
	Password		string
	PoolSize 		int
	Cluster 		[]string
}

func (lc *LoggerC) initLogger() error {
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

func (r *RedisC) initRedisCluster() error {
	//
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: r.Cluster,
	})

	err := client.Ping().Err()
	if err != nil {
		return err
	}
	redisClient = client

	return nil
}

func RedisClient() *redis.ClusterClient {
	return redisClient
}

