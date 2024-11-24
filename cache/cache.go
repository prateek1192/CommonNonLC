package cache

import (
	"errors"
	"sync"
)

func NewCache(capacity int) *Cache {

	return &Cache{
		data:      make(map[string]interface{}),
		cap:       capacity,
		evictList: nil,
		mutex:     sync.RWMutex{},
	}
}

func (c *Cache) Set(key string, value interface{}) {

	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.data) == c.cap {
		c.evict()
	}
	c.data[key] = value
	c.track(key)
}

func (c *Cache) evict() {
	if len(c.evictList) == 0 {
		return
	}
	evictKey := c.evictList[0]
	c.evictList = c.evictList[1:]
	delete(c.data, evictKey)
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mutex.RLock()
	value, exists := c.data[key]
	c.mutex.RUnlock()

	if !exists {
		return nil, errors.New("key not found")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.track(key)
	return value, nil

}

func (c *Cache) track(key string) {

	for i, v := range c.evictList {
		if v == key {
			c.evictList = append(c.evictList[:i], c.evictList[i+1:]...)
			break
		}
	}
	c.evictList = append(c.evictList, key)
}
