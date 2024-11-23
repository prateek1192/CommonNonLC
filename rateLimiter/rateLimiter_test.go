package rateLimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiter_Allow(t *testing.T) {

	limiter := NewRateLimiter(4, 3)

	for i := 0; i <= 50; i++ {
		time.Sleep(10 * time.Millisecond)
		if limiter.Allow() {
			fmt.Printf("Request %d: Allowed \n", i)
		} else {
			fmt.Printf("Request %d: Denied \n", i)
		}
	}
}
