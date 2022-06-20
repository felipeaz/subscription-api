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

// logger implements the logger interface
type logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLog returns an instance of Log
func NewLog() *logger {
	flags := log.LstdFlags | log.Lshortfile
	return &logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", flags),
		warnLogger:  log.New(os.Stdout, "[WARN] ", flags),
		errorLogger: log.New(os.Stdout, "[ERROR] ", flags),
	}
}

// Info log at INFO level
func (l *logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Warn log at WARN level
func (l *logger) Warn(v ...interface{}) {
	l.warnLogger.Println(v)
}

// Error log at ERROR level
func (l *logger) Error(v ...interface{}) {
	l.errorLogger.Println(v)
}
