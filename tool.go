package xlog

import (
	"runtime"
)

//获取运行函数的行数；需要了解函数调用栈相关知识
//skip调用深度  pc:program computer 程序计数器
func getLineInfo(skip int) (filename, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip) //runtime.Caller可以获取当前调用栈信息
	if ok {
		fun := runtime.FuncForPC(pc) //把pc传过来返回一个函数类型
		funcName = fun.Name()
	}
	filename = file
	lineNo = line
	return
}
