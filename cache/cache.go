package cache

import "errors"

func NewCache(capacity int) *Cache {

	return &Cache{
		data:      make(map[string]interface{}),
		cap:       capacity,
		evictList: nil,
	}
}

func (c *Cache) Set(key string, value interface{}) {

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
	if value, exists := c.data[key]; exists {
		c.track(key)
		return value, nil
	}
	return nil, errors.New("key not found")
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
