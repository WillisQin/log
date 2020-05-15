package log

import (
	"os"
)

//控制台类型日志的处理
//定义所需结构体
type XConsole struct {
	*XLogBase //继承基类
}

//构造函数，初始化
func NewXConsole(level int, module string) XLog {
	//XConsole实现了XLog接口，所以logger可以存储XConsole实例
	logger := &XConsole{}
	logger.XLogBase = &XLogBase{
		level:  level,
		module: module,
	}
	return logger
}

//定义实现每个日志级别的方法，即实现了XLog接口

func (c *XConsole) Init() (err error) {
	return nil
}

//日志格式：2018-10-26 01:17:40.222 DEBUG  use_service(模块) (file.go:LogDebug:332) this test txt
func (c *XConsole) LogDebug(format string, args ...interface{}) {

	//判断如果日志级别大于Debug，就不执行
	if c.level > XLogLevelDebug {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelDebug, c.module, format, args...)
	c.writeLog(os.Stdout, logData) //此时文件句柄为控制台，os.Stdout

}

func (c *XConsole) LogTrace(format string, args ...interface{}) {

	if c.level > XLogLevelTrace {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelTrace, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *XConsole) LogInfo(format string, args ...interface{}) {

	if c.level > XLogLevelInfo {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelInfo, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *XConsole) LogWarn(format string, args ...interface{}) {

	if c.level > XLogLevelWarn {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelWarn, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *XConsole) LogError(format string, args ...interface{}) {

	if c.level > XLogLevelError {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelError, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *XConsole) LogFatal(format string, args ...interface{}) {

	if c.level > XLogLevelFatal {
		return
	}

	logData := c.formatLoggerConsole(XLogLevelFatal, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *XConsole) SetLevel(level int) {
	c.level = level
}

func (c *XConsole) Close() {

}
