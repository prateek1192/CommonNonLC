package rateLimiter

import (
	"sync"
	"time"
)

func NewRateLimiter(rate, bucketCap int) *RateLimiter {

	return &RateLimiter{
		rate:      rate,
		bucketCap: bucketCap,
		tokens:    bucketCap,
		mutex:     sync.Mutex{},
		lastTime:  time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime).Seconds()
	rl.tokens = rl.tokens + int(elapsed*(float64(rl.rate)))
	if rl.tokens > rl.bucketCap {
		rl.tokens = rl.bucketCap
	}
	rl.lastTime = now

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}
