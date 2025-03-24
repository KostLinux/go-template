package logger

import (
	"fmt"
	"go-template/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func Setup(cfg *config.New) error {
	// Parse log level
	lvl, err := zapcore.ParseLevel(cfg.App.Log.Level)
	if err != nil {
		return fmt.Errorf("invalid log level %s: %w", cfg.App.Log.Level, err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(lvl),
		Development:      cfg.App.Env == "development",
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
		InitialFields: map[string]interface{}{
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
			"env":     cfg.App.Env,
		},
	}

	logger, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		return fmt.Errorf("failed to build logger: %w", err)
	}

	log = logger.Sugar()
	return nil
}
