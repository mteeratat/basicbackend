package customLog

import (
	"fmt"
	"io"
	"time"
)

const (
	LevelInfo = iota
	LevelWarn
	LevelError
)

var levelNames = map[int]string{
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
}

func getLevelName(level int) string {
	if name, ok := levelNames[level]; ok {
		return name
	}
	return "UNKNOWN"
}

type CustomLogger struct {
	level       int
	destination io.Writer
}

func NewCustomLogger(level int, destination io.Writer) *CustomLogger {
	return &CustomLogger{
		level:       level,
		destination: destination,
	}
}

func (l *CustomLogger) log(level int, msg string) {
	if level < l.level {
		return
	}

	logMessage := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Format(time.DateTime), getLevelName(level), msg)

	l.destination.Write([]byte(logMessage))
}

func (l *CustomLogger) Info(msg string) {
	l.log(LevelInfo, msg)
}
func (l *CustomLogger) Warn(msg string) {
	l.log(LevelWarn, msg)
}
func (l *CustomLogger) Error(msg string) {
	l.log(LevelError, msg)
}
