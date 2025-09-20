package tasks

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSafeCache(t *testing.T) {
	cache := NewCache()
	defer cache.Close()

	t.Run("Set and Get", func(t *testing.T) {
		cache.Set("key1", "value1", 5*time.Second)
		value, ok := cache.Get("key1")
		if !ok || value != "value1" {
			t.Errorf("expected value1, got %v", value)
		}
	})

	t.Run("Get expired key", func(t *testing.T) {
		cache.Set("key2", "value2", 1*time.Millisecond)
		time.Sleep(2 * time.Millisecond)
		_, ok := cache.Get("key2")
		if ok {
			t.Error("expected key2 to be expired")
		}
	})

	t.Run("Delete key", func(t *testing.T) {
		cache.Set("key3", "value3", 5*time.Second)
		cache.Delete("key3")
		_, ok := cache.Get("key3")
		if ok {
			t.Error("expected key3 to be deleted")
		}
	})

	t.Run("Concurrent access no deadlocks", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				key := "key"
				cache.Set(key, i, time.Duration(rand.Intn(1000))*time.Millisecond)
				_, _ = cache.Get(key)
			}(i)
		}
		wg.Wait()
	})
}
