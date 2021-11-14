package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		trans, found := global.Ut.GetTranslator(locale)
		if found {
			c.Set("trans", trans)
		} else {
			c.Set("trans", "en")
		}
		c.Next()
	}
}
