package cache

type Cache struct {
	data      map[string]interface{}
	cap       int
	evictList []string
}
