package limiter

import (
	"sync"
	"time"
)

var _ Limiter = &SlidingWindowLimiter{}

type SlidingWindowLimiter struct {
	windowSize   time.Duration // 窗口大小
	maxRequests  int           // 最大请求数
	requests     []time.Time   // 窗口内的请求时间
	requestsLock sync.Mutex    // 请求锁
}

func NewSlidingWindowLimiter(windowSize time.Duration, maxRequests int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		windowSize:  windowSize,
		maxRequests: maxRequests,
		requests:    make([]time.Time, 0),
	}
}

func (limiter *SlidingWindowLimiter) Allow() bool {
	limiter.requestsLock.Lock()
	defer limiter.requestsLock.Unlock()

	// 移除过期的请求
	currentTime := time.Now()
	for len(limiter.requests) > 0 && currentTime.Sub(limiter.requests[0]) > limiter.windowSize {
		limiter.requests = limiter.requests[1:]
	}

	// 检查请求数是否超过阈值
	if len(limiter.requests) >= limiter.maxRequests {
		return false
	}

	limiter.requests = append(limiter.requests, currentTime)
	return true
}
