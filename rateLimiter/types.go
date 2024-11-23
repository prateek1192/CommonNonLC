package rateLimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate      int        // Max Tokens added per second
	bucketCap int        // Maximum tokens in the bucket
	tokens    int        // Current tokens in the bucket
	mutex     sync.Mutex // Mutex to handle concurrency
	lastTime  time.Time  // Last time tokens were added
}
