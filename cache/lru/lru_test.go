package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {

	cache := NewLRUCache(3)

	cache.Put("1", "1")
	cache.Put("2", "2")
	cache.Put("3", "3")
	cache.Put("4", "4")
	cache.Get("2")
	value := cache.Get("1")
	cache.Put("1", "1")

	fmt.Println("value", value)
	fmt.Println("size: ",cache.doubleList.Size)
	cache.doubleList.Display() // 1 2 4
}
