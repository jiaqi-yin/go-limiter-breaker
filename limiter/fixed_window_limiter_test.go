package limiter_test

import (
	"fmt"
	"testing"
	"time"

	"go-limiter-breaker/limiter"
)

func TestFixedWindowLimiter(t *testing.T) {
	// 每秒最多允许3个请求
	limiter := limiter.NewFixedWindowLimiter(time.Second, 3)

	for i := 0; i < 15; i++ {
		now := time.Now().Format("15:04:05")

		if limiter.Allow() {
			fmt.Println(now + " 请求通过")
		} else {
			fmt.Println(now + " 请求被限流")
		}

		time.Sleep(100 * time.Millisecond)
	}
}
