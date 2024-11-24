package cache

import "sync"

type Cache struct {
	data      map[string]interface{}
	cap       int
	evictList []string
	mutex     sync.RWMutex
}
