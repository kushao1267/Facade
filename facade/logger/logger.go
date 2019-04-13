package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// level logger
var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

const (
	// Log level
	INFO  = 1
	DEBUG = 2
	ERROR = 3
	FATAL = 4
)

type Logger struct{}

// exported json logger
var JsonLogger Logger

func init() {
	AppLogPath := os.Getenv("log_path")
	if len(AppLogPath) == 0{
		AppLogPath = "./"
	}

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
	Fatal = log.New(file, "FATAL ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 日志内容,map类型
type Content map[string]interface{}

// print 基础log方法
func (l *Logger) print(level int, msg string, keyValue *Content) {
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
	case FATAL:
		Fatal.Println(printMsg)
	default:
		Info.Println(printMsg)
	}
	fmt.Println(printMsg)
}

func (l *Logger) save(level int, msg string, keyValue ...interface{}) {

	if len(keyValue) > 0 {
		switch v := keyValue[0].(type) {
		case *Content: // 字典类型
			l.print(level, msg, v)
		case string:
			l.print(level, msg+v, nil)
		case int:
			l.print(level, msg+string(v), nil)
		case []string:
			l.print(level, msg+strings.Join(v, ", "), nil)
		}
		return
	}
	l.print(level, msg, nil)
}

// Info logger
func (l *Logger) Info(msg interface{}, keyValue ...interface{}) {
	switch m := msg.(type) {
	case string:
		l.save(INFO, m, keyValue)
	case error:
		l.save(INFO, m.Error(), keyValue)
	}
}

// Debug logger
func (l *Logger) Debug(msg interface{}, keyValue ...interface{}) {
	switch m := msg.(type) {
	case string:
		l.save(DEBUG, m, keyValue)
	case error:
		l.save(DEBUG, m.Error(), keyValue)
	}
}

// Error logger
func (l *Logger) Error(msg interface{}, keyValue ...interface{}) {
	switch m := msg.(type) {
	case string:
		l.save(ERROR, m, keyValue)
	case error:
		l.save(ERROR, m.Error(), keyValue)
	}
}

// Fatal logger
func (l *Logger) Fatal(msg interface{}, keyValue ...interface{}) {
	switch m := msg.(type) {
	case string:
		l.save(FATAL, m, keyValue)
	case error:
		l.save(FATAL, m.Error(), keyValue)
	}
}
