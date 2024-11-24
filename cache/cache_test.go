package cache

import (
	"strconv"
	"sync"
	"testing"
)

func TestCache_Set_Get(t *testing.T) {
	c := NewCache(5)
	c.Set("3", 5)
	ans, _ := c.Get("3")
	if ans.(int) != 5 {
		t.Errorf("Unexpected Value")
	}
}

func TestCache_GetMostRecent(t *testing.T) {
	c := NewCache(5)
	c.Set("3", 5)
	c.Set("3", 7)
	ans, _ := c.Get("3")
	if ans.(int) != 7 {
		t.Errorf("Unexpected Value")
	}
}

func TestCache_ConfirmEviction(t *testing.T) {
	c := NewCache(3)
	c.Set("3", 5)
	c.Set("3", 7)
	c.Set("4", 8)
	c.Set("5", 9)
	c.Set("6", 10)

	_, err := c.Get("3")

	if err == nil {
		t.Errorf("Eviction not working")
	}
}

func TestCacheConcurrency(t *testing.T) {

	c := NewCache(5)
	var wg sync.WaitGroup

	numOperations := 100
	numGoroutines := 10

	setWorker := func(workerId int) {
		defer wg.Done()
		for i := 0; i < numOperations; i++ {
			key := "key_" + strconv.Itoa(workerId) + strconv.Itoa(i)
			c.Set(key, i)
		}
	}

	getWorker := func(workerId int) {
		defer wg.Done()
		for i := 0; i < numOperations; i++ {
			key := "key_" + strconv.Itoa(workerId) + strconv.Itoa(i)
			c.Get(key)
		}
	}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go setWorker(i)
		wg.Add(1)
		go getWorker(i)
	}

	wg.Wait()

	if len(c.data) > c.cap {
		t.Errorf("Cache exceeded capacity: got %d items, expected at most %d", len(c.data), c.cap)
	}

	t.Log("All operations completed successfully.")

}
