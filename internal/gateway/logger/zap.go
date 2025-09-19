package logger

import (
	"game-server-golang/internal/config"
	"game-server-golang/internal/gateway"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	zap *zap.Logger
}

// NewZapLogger creates a new Zap logger with the given configuration
func NewZapLogger(cfg config.LoggingConfig) gateway.Logger {
	// Configure log level
	var level zapcore.Level
	switch cfg.Level {
	case config.LogLevelDebug:
		level = zap.DebugLevel
	case config.LogLevelInfo:
		level = zap.InfoLevel
	case config.LogLevelWarn:
		level = zap.WarnLevel
	case config.LogLevelError:
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// Create logger configuration
	zapCfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         cfg.Format,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if cfg.DisableCaller {
		zapCfg.EncoderConfig.CallerKey = "caller"
		zapCfg.DisableCaller = true
	}

	// Customize time format
	zapCfg.EncoderConfig.TimeKey = "timestamp"
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ := zapCfg.Build(
		zap.AddCallerSkip(1), // Skip the wrapper calls
	)

	return &ZapLogger{
		zap: logger,
	}
}

func (l *ZapLogger) Debug(msg string) {
	l.zap.Debug(msg)
}

func (l *ZapLogger) Info(msg string) {
	l.zap.Info(msg)
}

func (l *ZapLogger) Warn(msg string) {
	l.zap.Warn(msg)
}

func (l *ZapLogger) Error(msg string) {
	l.zap.Error(msg)
}

func (l *ZapLogger) WithField(key string, value any) gateway.Logger {
	return &ZapLogger{
		zap: l.zap.With(zap.Any(key, value)),
	}
}

func (l *ZapLogger) WithFields(fields map[string]any) gateway.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return &ZapLogger{
		zap: l.zap.With(zapFields...),
	}
}
