package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ecode := errcode.Success
		token := c.GetHeader("token")
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				ecode = errcode.UnauthorizedTokenError
			} else if time.Now().Unix() > claims.ExpiresAt {
				ecode = errcode.UnauthorizedTokenTimeout
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
