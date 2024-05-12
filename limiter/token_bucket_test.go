package limiter_test

import (
	"fmt"
	"testing"
	"time"

	"go-limiter-breaker/limiter"
)

func TestTokenBucket(t *testing.T) {
	tokenBucket := limiter.NewTokenBucket(2.0, 3.0)

	for i := 1; i <= 20; i++ {
		now := time.Now().Format("15:04:05")

		if tokenBucket.Allow() {
			fmt.Printf(now+"  第 %d 个请求通过\n", i)
		} else { // 如果不能移除一个令牌，请求被拒绝。
			fmt.Printf(now+"  第 %d 个请求被限流\n", i)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
