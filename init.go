package AdvancedLogging

import (
	"fmt"
	"time"
)

type ALog struct {
	Level     string `json:"level"`
	Timestamp any    `json:"timestamp"`
	Log       Log    `json:"log"`
}

type Log struct {
	Sender  string `json:"sender"`
	Message any    `json:"message"`
}

func Info(log Log) ALog {
	alog := ALog{
		Level:     "INFO",
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Log:       log,
	}

	fmt.Println(fmt.Sprintf("%sINFO  %v %s : %v%s", InfoColor, alog.Timestamp, log.Sender, log.Message, Reset))
	return alog
}

func Warn(log Log) ALog {
	alog := ALog{
		Level:     "WARN",
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Log:       log,
	}

	fmt.Println(fmt.Sprintf("%sWARN  %v %s : %v%s", WarnColor, alog.Timestamp, log.Sender, log.Message, Reset))
	return alog
}

func Error(log Log) ALog {
	alog := ALog{
		Level:     "ERROR",
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Log:       log,
	}

	fmt.Println(fmt.Sprintf("%sERROR %v %s : %v%s", ErrorColor, alog.Timestamp, log.Sender, log.Message, Reset))
	return alog
}

func Debug(log Log) ALog {
	alog := ALog{
		Level:     "DEBUG",
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Log:       log,
	}

	fmt.Println(fmt.Sprintf("%sDEBUG %v %s : %v%s", DebugColor, alog.Timestamp, log.Sender, log.Message, Reset))
	return alog
}
