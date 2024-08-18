package logger

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string, err error)
	Debug(msg string)
	InfoMsg(msg string, fields map[string]interface{})
	Print(msg string, args ...interface{})
}
