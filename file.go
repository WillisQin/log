package xlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//文件类型日志的处理

//定义所需结构体
type XFile struct {
	filename  string   //文件名称
	File      *os.File //文件实例
	*XLogBase          //继承基类

	LogChan chan *LogData //写文件操作异步处理，通过管道实现
	wg      *sync.WaitGroup

	curHour string //日志文件切割，根据日期切割
}

//构造函数，初始化
func NewXFile(level int, filename, module string) XLog {
	//XFile实现了XLog接口，所以logger可以存储XFile实例
	logger := &XFile{
		filename: filename,
	}
	logger.XLogBase = &XLogBase{
		level:  level,
		module: module,
	}

	logger.curHour = time.Now().Format("2006010215") //日志切割，获取初始化文件是的小时时间
	logger.wg = &sync.WaitGroup{}
	logger.LogChan = make(chan *LogData, 100000)
	logger.wg.Add(1)
	go logger.syncLog()
	return logger
}

//初始化文件
func (c *XFile) Init() (err error) {
	c.File, err = os.OpenFile(c.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	return
}

//写chan
func (c *XFile) writeToChan(level int, module string, format string, args ...interface{}) {

	logData := c.formatLogger(level, module, format, args...)
	//c.writeLog(c.file, logData)
	//异步写入文件，chan实现。当写入速度大于取出速度时chan会堵塞，导致程序崩溃；通过select写入chan，防止堵塞
	select {
	case c.LogChan <- logData:
	default:
	}
}

//goroutine取chan，写入文件
func (c *XFile) syncLog() {

	for data := range c.LogChan {
		c.splitLog()
		c.writeLog(c.File, data)
	}

	c.wg.Done()
}

//切割日志
func (c *XFile) splitLog() {
	now := time.Now()
	hour := now.Format("2006010215")
	if hour == c.curHour {
		return
	}

	c.curHour = hour //切割文件后重新定义curHour变量
	c.File.Sync()
	c.File.Close()
	var newFilename string
	if now.Hour() == 00 {
		newFilename = fmt.Sprintf("%s.%s.log", c.filename, fmt.Sprintf("%4d-%02d-%02d-%02d", now.Year(), now.Month(), now.Day()-1, 23))
	} else {
		newFilename = fmt.Sprintf("%s.%s.log", c.filename, fmt.Sprintf("%4d-%02d-%02d-%02d", now.Year(), now.Month(), now.Day(), now.Hour()-1))
	}
	os.Rename(c.filename, newFilename)
	c.Init()
}

//定义实现每个日志级别的方法,即实现了XLog接口
func (c *XFile) LogDebug(format string, args ...interface{}) {

	//判断如果日志级别大于Debug，就不执行
	if c.level > XLogLevelDebug {
		return
	}

	c.writeToChan(XLogLevelDebug, c.module, format, args...)

}

func (c *XFile) LogTrace(format string, args ...interface{}) {
	if c.level > XLogLevelTrace {
		return
	}

	c.writeToChan(XLogLevelTrace, c.module, format, args...)
}

func (c *XFile) LogInfo(format string, args ...interface{}) {
	if c.level > XLogLevelInfo {
		return
	}

	c.writeToChan(XLogLevelInfo, c.module, format, args...)
}

func (c *XFile) LogWarn(format string, args ...interface{}) {
	if c.level > XLogLevelWarn {
		return
	}

	c.writeToChan(XLogLevelWarn, c.module, format, args...)
}

func (c *XFile) LogError(format string, args ...interface{}) {
	if c.level > XLogLevelError {
		return
	}

	c.writeToChan(XLogLevelError, c.module, format, args...)
}

func (c *XFile) LogFatal(format string, args ...interface{}) {
	if c.level > XLogLevelFatal {
		return
	}

	c.writeToChan(XLogLevelFatal, c.module, format, args...)
}

func (c *XFile) SetLevel(level int) {
	c.level = level
}

//关闭文件
func (c *XFile) Close() {

	//wait之前关闭for 管道
	if c.LogChan != nil {
		close(c.LogChan)
	}
	c.wg.Wait()

	if c.File != nil {
		c.File.Sync() //关闭文件之前把缓冲区内容刷到磁盘
		c.File.Close()
	}
}
