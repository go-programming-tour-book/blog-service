package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"github.com/go-programming-tour-book/blog-service/global"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		switch locale {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(global.ValidatorV9.Validate, trans)
			break
		case "en":
			_ = en_translations.RegisterDefaultTranslations(global.ValidatorV9.Validate, trans)
			break
		default:
			_ = zh_translations.RegisterDefaultTranslations(global.ValidatorV9.Validate, trans)
			break
		}

		c.Set("trans", trans)
		c.Next()
	}
}
