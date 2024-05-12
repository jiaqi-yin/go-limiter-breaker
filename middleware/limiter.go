package middleware

import (
	"go-limiter-breaker/limiter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Limiter(l limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !l.Allow() {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "0 tokens available, try later",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
