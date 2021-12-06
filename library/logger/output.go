package logger

import (
	"os"
)

// 获取输出到文件 到 io 对象
func GetOutputFile(path string) (file *os.File, err error) {
	return os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
}