package limiter

import "time"

var _ Limiter = &LeakyBucket{}

type LeakyBucket struct {
	rate       float64 // 漏桶速率，单位请求数/秒
	capacity   int     // 漏桶容量，最多可存储请求数
	water      int     // 当前水量，表示当前漏桶中的请求数
	lastLeakMs int64   // 上次漏水的时间戳，单位秒
}

func NewLeakyBucket(rate float64, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:       rate,
		capacity:   capacity,
		water:      0,
		lastLeakMs: time.Now().Unix(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	now := time.Now().Unix()
	elapsed := now - lb.lastLeakMs

	// 漏水，根据时间间隔计算漏掉的水量
	leakAmount := int(float64(elapsed) / 1000 * lb.rate)
	if leakAmount > 0 {
		if leakAmount > lb.water {
			lb.water = 0
		} else {
			lb.water -= leakAmount
		}
	}

	// 判断当前水量是否超过容量
	if lb.water > lb.capacity {
		lb.water-- // 如果超过容量，减去刚刚增加的水量
		return false
	}

	// 增加水量
	lb.water++

	lb.lastLeakMs = now
	return true
}
