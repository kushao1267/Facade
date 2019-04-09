package logger

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

// level logger
var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

const (
	AppLogPath = "/data/app/log/"
	// Log level
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

type Logger struct{}

// exported json logger
var JsonLogger Logger

func init() {
	// 指定log文件
	os.MkdirAll(AppLogPath, os.ModePerm) // mkdir -p
	logPath := path.Join(AppLogPath, "app.log")
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// 初始化logger
	Info = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(file, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 日志内容,map类型
type Content map[string]interface{}

// Save 基础log方法
func (l *Logger) Save(level int, msg interface{}, keyValue *Content) {
	// 初始化return map
	if keyValue == nil {
		keyValue = &Content{}
	}
	// 返回message,可是string或者map类型
	(*keyValue)["message"] = msg
	jsonMsg, _ := json.Marshal(keyValue)
	printMsg := string(jsonMsg)
	switch level {
	case INFO:
		Info.Println(printMsg)
	case DEBUG:
		Debug.Println(printMsg)
	case ERROR:
		Error.Println(printMsg)
	default:
		Info.Println(printMsg)
	}
}

// Info logger
func (l *Logger) Info(msg interface{}, keyValue *Content) {
	l.Save(INFO, msg, keyValue)
}

// Debug logger
func (l *Logger) Debug(msg interface{}, keyValue *Content) {
	l.Save(DEBUG, msg, keyValue)
}

// Error logger
func (l *Logger) Error(msg interface{}, keyValue *Content) {
	l.Save(ERROR, msg, keyValue)
}
