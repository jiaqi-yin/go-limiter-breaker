package limiter

import (
	"sync"
	"time"
)

var _ Limiter = &FixedWindowLimiter{}

type FixedWindowLimiter struct {
	windowSize  time.Duration // 窗口大小
	maxRequests int           // 最大请求数
	requests    int           // 当前窗口内的请求数
	lastReset   int64         // 上次窗口重置时间（秒级时间戳）
	resetMutex  sync.Mutex    // 重置锁
}

func NewFixedWindowLimiter(windowSize time.Duration, maxRequests int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		windowSize:  windowSize,
		maxRequests: maxRequests,
		lastReset:   time.Now().Unix(),
	}
}

func (limiter *FixedWindowLimiter) Allow() bool {
	limiter.resetMutex.Lock()
	defer limiter.resetMutex.Unlock()

	// 检查是否需要重置窗口
	if time.Now().Unix()-limiter.lastReset >= int64(limiter.windowSize.Seconds()) {
		limiter.requests = 0
		limiter.lastReset = time.Now().Unix()
	}

	// 检查请求数是否超过阈值
	if limiter.requests >= limiter.maxRequests {
		return false
	}

	limiter.requests++
	return true
}
