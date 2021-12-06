package validator

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"sync"
)

type validate struct {
	*validator.Validate
	translators map[string]ut.Translator
}

var (
	instance *validate
	once sync.Once
)

// 单例模式
func GetInstance() *validate {
	once.Do(func() {
		// gin 框架 Validator 绑定
		ins, _ := binding.Validator.Engine().(*validator.Validate)
		instance = &validate{
			Validate: ins,
			translators: make(map[string]ut.Translator, 0),
		}
		// 翻译
		instance.translate()
		// tag 注册 自定义注册验证规则
		instance.registerTag()
	})
	return instance
}