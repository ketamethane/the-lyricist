package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *log.Logger
}

func NewLogrusLogger() Logger {

	logger := log.New()
	filePath := "c/error"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}

	logger.SetOutput(file)
	logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	logger.SetLevel(log.DebugLevel)

	return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *LogrusLogger) InfoMsg(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(msg)

}

func (l *LogrusLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *LogrusLogger) Error(msg string, err error) {
	l.logger.Errorf("%s: %v", msg, err)
}

func (l *LogrusLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

// Print implements the Print method for LogrusLogger
func (l *LogrusLogger) Print(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}
