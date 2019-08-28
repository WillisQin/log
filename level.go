package xlog

//日志级别常量定义 1 2 ...以此类推用ioto自动设置常量
const (
	XLogLevelDebug = iota
	XLogLevelTrace
	XLogLevelInfo
	XLogLevelWarn
	XLogLevelError
	XLogLevelFatal
)

//日志类型常量定义
const (
	XLogTypeFile = iota
	XLogTypeConsole
)

//根据int类型日志级别转换成对应的字符串类型
func getLevelStr(level int) string {
	switch level {
	case XLogLevelDebug:
		return "DEBUG"
	case XLogLevelTrace:
		return "TRACE"
	case XLogLevelInfo:
		return "INFO"
	case XLogLevelWarn:
		return "WARN"
	case XLogLevelError:
		return "ERROR"
	case XLogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}

}
