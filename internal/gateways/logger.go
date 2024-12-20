package gateways

type Logger interface {
	Info(msg string)
	WithField(key string, value interface{}) Logger
}
