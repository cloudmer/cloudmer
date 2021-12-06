package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 翻译
func (ins *validate) translate()  {
	// 注册英文错误提示器
	_en := en.New()
	ins.translators["en"], _ =  ut.New(_en, _en).GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(ins.Validate, ins.translators["en"])

	// 注册中文错误提示器
	_zh := zh.New()
	ins.translators["zh"], _ = ut.New(_zh, _zh).GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(ins.Validate, ins.translators["zh"])
}