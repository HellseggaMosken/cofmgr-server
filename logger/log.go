package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var lg = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

// Init Init the logger
func Init() {
	level := logrus.InfoLevel
	switch os.Getenv("LOG_LEVEL") {
	case "error":
		level = logrus.ErrorLevel
	case "warning":
		level = logrus.WarnLevel
	case "info":
		level = logrus.InfoLevel
	case "debug":
		level = logrus.DebugLevel
	}
	lg = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}
}

func Debug(msg ...interface{}) {
	lg.Debugln(msg...)
}

func Info(msg ...interface{}) {
	lg.Infoln(msg...)
}

func Warning(msg ...interface{}) {
	lg.Warningln(msg...)
}

func Error(msg ...interface{}) {
	lg.Errorln(msg...)
}

func Fatal(msg ...interface{}) {
	lg.Fatalln(msg...)
	os.Exit(0)
}

func Panic(msg ...interface{}) {
	lg.Panicln(msg...)
}

func Logger() *logrus.Logger {
	return lg
}

func LoggerLevel() logrus.Level {
	return lg.Level
}
