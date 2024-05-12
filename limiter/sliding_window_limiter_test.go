package limiter_test

import (
	"fmt"
	"testing"
	"time"

	"go-limiter-breaker/limiter"
)

func TestSlidingWindowLimiter(t *testing.T) {
	// 每秒最多允许4个请求
	limiter := limiter.NewSlidingWindowLimiter(500*time.Millisecond, 2)

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
