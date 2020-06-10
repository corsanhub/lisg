package logging

import (
	"fmt"
	"time"

	"corsanhub.com/lisg/corsan/util"
)

var logFile = "module.log"

// var GlobalLogLevel = DEBUG

var GlobalLogLevel = INFO
var GlobalLogType = BOTH

type Logging interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type LogType int
type LogLevel int

const (
	BOTH LogType = iota
	SYSOUT
	FILE
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	Name string
	Type LogType
}

func logLevelStr(level LogLevel) string {
	result := ""
	switch level {
	case DEBUG:
		result = "DEBUG"
	case INFO:
		result = "INFO "
	case WARN:
		result = "WARN "
	case ERROR:
		result = "ERROR"
	case FATAL:
		result = "FATAL"
	}
	return result
}

func (logger Logger) printLog(level string, str string) {
	currentTime := time.Now()
	dateStr := currentTime.Format("2006-01-02 15:04:05.000")

	logStr := dateStr + " [" + level + "] - " + util.TraceStr(2) + ":" + str

	switch logger.Type {
	case SYSOUT:
		fmt.Println(logStr)
	case FILE:
		util.AppendLineToFile(logFile, logStr)
	case BOTH:
		fmt.Println(logStr)
		util.AppendLineToFile(logFile, logStr)
	}
}

func (logger Logger) Log(str string, desiredLevel LogLevel) {
	//fmt.Printf("desiredLevel: %v, GlobalLogLevel: %v", desiredLevel, GlobalLogLevel)
	if desiredLevel >= GlobalLogLevel {
		logger.printLog(logLevelStr(desiredLevel), str)
	}
}

func (logger Logger) Debug(str string) {
	logger.Log(str, DEBUG)
}

func (logger Logger) Info(str string) {
	logger.Log(str, INFO)
}

func (logger Logger) Warn(str string) {
	logger.Log(str, WARN)
}

func (logger Logger) Error(str string) {
	logger.Log(str, ERROR)
}

func (logger Logger) fatal(str string) {
	logger.Log(str, FATAL)
}
