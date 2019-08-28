package xlog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//定义基类，file，console都公用的类

//简化代码，把日志格式化代码单独封装
type LogData struct {
	timeStr  string
	levelStr string
	module   string
	filename string
	funcName string
	lineNo   int
	data     string
}

type XLogBase struct {
	level  int    //日志级别
	module string //模块名
}

func (l *XLogBase) writeLog(file *os.File, logData *LogData) {
	fmt.Fprintf(file, "%s %s %s [%s:%s:%d]  %s\n",
		logData.timeStr, logData.levelStr, logData.module,
		logData.filename, logData.funcName, logData.lineNo, logData.data)
}

//文件日志操作多了一层管道操作，获取行号时GetLineInfo深度比命令行日志多1层，把深度单独传参
func (l *XLogBase) formatLogger(level int, module string, format string, args ...interface{}) *LogData {

	//获取当前时间
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000")
	//获取日志级别
	levelStr := getLevelStr(level)
	//获取模块名
	filename, funcName, lineNo := getLineInfo(5)
	filename = filepath.Base(filename)
	data := fmt.Sprintf(format, args...)

	return &LogData{
		timeStr:  timeStr,
		levelStr: levelStr,
		module:   module,
		filename: filename,
		funcName: funcName,
		lineNo:   lineNo,
		data:     data,
	}
}

//文件日志操作多了一层管道操作，获取行号时GetLineInfo深度比命令行日志多1层，把深度单独传参
func (l *XLogBase) formatLoggerConsole(level int, module string, format string, args ...interface{}) *LogData {

	//获取当前时间
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000")
	//获取日志级别
	levelStr := getLevelStr(level)
	//获取模块名
	filename, funcName, lineNo := getLineInfo(4)
	filename = filepath.Base(filename)
	data := fmt.Sprintf(format, args...)

	return &LogData{
		timeStr:  timeStr,
		levelStr: levelStr,
		module:   module,
		filename: filename,
		funcName: funcName,
		lineNo:   lineNo,
		data:     data,
	}

}
