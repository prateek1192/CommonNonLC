package cache

import (
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
