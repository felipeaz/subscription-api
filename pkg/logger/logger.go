package logger

import (
	"log"
	"os"
)

const (
	logFile = "%s-log"
)

type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

// Log implements the logger interface
type Log struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLog returns an instance of Log
func NewLog() *Log {
	flags := log.LstdFlags | log.Lshortfile
	return &Log{
		infoLogger:  log.New(os.Stdout, "[INFO] ", flags),
		warnLogger:  log.New(os.Stdout, "[WARN] ", flags),
		errorLogger: log.New(os.Stdout, "[ERROR] ", flags),
	}
}

// Info log at INFO level
func (l *Log) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Warn log at WARN level
func (l *Log) Warn(v ...interface{}) {
	l.warnLogger.Println(v)
}

// Error log at ERROR level
func (l *Log) Error(v ...interface{}) {
	l.errorLogger.Println(v)
}
