package logger

import (
	"sample_go_grpc_server/utils/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service interface {
	NewPrefix(prefix string) *CtxLogger
}

type loggerService struct {
	*zap.Logger
}

func (ls *loggerService) NewPrefix(prefix string) *CtxLogger {
	return &CtxLogger{ls.Logger.Named(prefix)}
}

func NewLoggerService(conf config.Config) Service {
	cfg := newConfig(conf.Logger)

	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return NewLoggerServiceWithZap(zapLogger)
}

func NewLoggerServiceWithZap(zl *zap.Logger) Service {
	return &loggerService{zl}
}

func utcTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z0700"))
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger_name",
		CallerKey:      "caller_file",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     utcTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newConfig(conf config.Logger) zap.Config {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	if conf.IsDevMode {
		cfg.Development = true
		cfg.Encoding = "console"
	}

	switch conf.Level {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warning":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	return cfg
}
