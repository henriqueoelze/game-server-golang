package gateway

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger
}
