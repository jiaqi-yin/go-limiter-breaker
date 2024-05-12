package main

import (
	"errors"
	"go-limiter-breaker/breaker"
	"go-limiter-breaker/limiter"
	"go-limiter-breaker/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Apply limiter globally.
	// r.Use(middleware.Limiter(limiter.NewTokenBucket(3.0, 5.0)))
	r.GET(
		"/ping/limiter",
		middleware.Limiter(limiter.NewTokenBucket(3.0, 5.0)),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong limiter",
			})
		},
	)

	b := breaker.NewBreaker(4, 4, 4, 15*time.Second)
	r.GET("/ping/breaker", func(c *gin.Context) {
		err := b.Exec(func() error {
			value, _ := c.GetQuery("value")
			if value == "0" {
				return errors.New("error")
			}
			return nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong breaker",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
