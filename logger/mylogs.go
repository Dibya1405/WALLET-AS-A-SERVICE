package logger

import (
	"log"
	"os"
)

type MainLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *MainLogger) Info(s ...interface{}) {
	l.infoLogger.Println(s)
}

func (l *MainLogger) Warn(s ...interface{}) {
	l.warnLogger.Println(s)
}

func (l *MainLogger) Error(s ...interface{}) {
	l.errorLogger.Println(s)
}

func GetMyLogger() *MainLogger {
	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARN: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)

	MyLogger := MainLogger{infoLogger, warnLogger, errorLogger}
	return &MyLogger
}
