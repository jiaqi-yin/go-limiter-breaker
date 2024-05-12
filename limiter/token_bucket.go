package limiter

import (
	"sync"
	"time"
)

var _ Limiter = &TokenBucket{}

// TokenBucket 表示一个令牌桶。
type TokenBucket struct {
	rate       float64    // 令牌添加到桶中的速率。
	capacity   float64    // 桶的最大容量。
	tokens     float64    // 当前桶中的令牌数量。
	lastUpdate time.Time  // 上次更新令牌数量的时间。
	mu         sync.Mutex // 互斥锁，确保线程安全。
}

// NewTokenBucket 创建一个新的令牌桶，给定令牌添加速率和桶的容量。
func NewTokenBucket(rate float64, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity, // 初始时，桶是满的。
		lastUpdate: time.Now(),
	}
}

// Allow 检查是否可以从桶中移除一个令牌。如果可以，它移除一个令牌并返回 true。
// 如果不可以，它返回 false。
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// 根据经过的时间计算要添加的令牌数量。
	now := time.Now()
	elapsed := now.Sub(tb.lastUpdate).Seconds()
	tokensToAdd := elapsed * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity // 确保令牌数量不超过桶的容量。
	}

	if tb.tokens >= 1.0 {
		tb.tokens--
		tb.lastUpdate = now
		return true
	}

	return false
}
