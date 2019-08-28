package xlog

//定义全局变量；防止用户不初始化logger(接口类型变量不初始化会pinic)，给默认值为console;日志初始化之前可以用日志库追踪问题
var logger XLog = newXLog(XLogTypeConsole, XLogLevelDebug, "", "default")

//定义XLog接口
type XLog interface {
	//接口方法和签名,用户可以传入字符串和其他可变参数
	Init() error //初始化文件方法
	LogDebug(fmt string, args ...interface{})
	LogTrace(fmt string, args ...interface{})
	LogInfo(fmt string, args ...interface{})
	LogWarn(fmt string, args ...interface{})
	LogError(fmt string, args ...interface{})
	LogFatal(fmt string, args ...interface{})
	Close() //关闭文件
	//可以让用户设置日志级别的方法，常量在level.go中定义
	SetLevel(level int)
}

//go中没有构造函数，需要用户自己定义
/*
定义一个写文件的函数：
logType:日志类型 -常量
level：日志级别  -常量
filename：需要打印到文件时的文件路径
module:模块名
*/
func newXLog(logType, level int, filename, module string) XLog {

	var logger XLog //定义一个借口类型变量logger
	//判断用户输入的日志类型
	switch logType {
	//文件类型日志使用file.go类处理
	case XLogTypeFile:
		logger = NewXFile(level, filename, module) //调用XFile的初始化构造函数
	//控制台类型日志使用console.go类处理
	case XLogTypeConsole:
		logger = NewXConsole(level, module) //调用XConsole的初始化构造函数
	default:
		logger = NewXFile(level, filename, module)
	}
	return logger //返回一个实例
}

func Init(logType, level int, filename, module string) error {
	logger = newXLog(logType, level, filename, module)
	return logger.Init()
}

//方便用户使用，封装函数
func LogDebug(fmt string, args ...interface{}) {
	logger.LogDebug(fmt, args...)
}
func LogTrace(fmt string, args ...interface{}) {
	logger.LogTrace(fmt, args...)
}
func LogInfo(fmt string, args ...interface{}) {
	logger.LogInfo(fmt, args...)
}
func LogWarn(fmt string, args ...interface{}) {
	logger.LogWarn(fmt, args...)
}
func LogError(fmt string, args ...interface{}) {
	logger.LogError(fmt, args...)
}
func LogFatal(fmt string, args ...interface{}) {
	logger.LogFatal(fmt, args...)
}
func Close() {
	logger.Close()
}

//可以让用户设置日志级别的方法，常量在level.go中定义
func SetLevel(level int) {
	logger.SetLevel(level)
}
