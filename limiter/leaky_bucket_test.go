package limiter_test

import (
	"fmt"
	"testing"
	"time"

	"go-limiter-breaker/limiter"
)

func TestLeakyBucket(t *testing.T) {
	// 创建一个漏桶，速率为每秒处理3个请求，容量为4个请求
	leakyBucket := limiter.NewLeakyBucket(3, 4)

	// 模拟请求
	for i := 1; i <= 15; i++ {
		now := time.Now().Format("15:04:05")

		if leakyBucket.Allow() {
			fmt.Printf(now+"  第 %d 个请求通过\n", i)
		} else {
			fmt.Printf(now+"  第 %d 个请求被限流\n", i)
		}

		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}
}
