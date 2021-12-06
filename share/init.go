package share

import (
	"cloudmer/library/logger"
	"cloudmer/library/validator"
)

func init()  {
	// console 日志输出
	Log = logger.DefaultConfig().Build()
	// validator v10 验证 内部初始化
	validator.GetInstance()
}