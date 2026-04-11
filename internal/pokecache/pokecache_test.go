package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	interval := 5 * time.Second
	testCases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://www.test.com",
			value: []byte("this is a test"),
		},
		{
			key:   "https://www.test2.com",
			value: []byte("this is another test"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCacheAdd %d", i+1), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.value)
			cacheEntry, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected cacheEntry with key: %s", tc.key)
			}
			if string(cacheEntry) != string(tc.value) {
				t.Errorf("expected cacheEntry: %s, but received: %s", string(tc.value), string(cacheEntry))
			}
		})
	}
}

func TestCleanUp(t *testing.T) {
	t.Run("Testing that cache is cleared after a given interval", func(t *testing.T) {
		interval := time.Millisecond * 5
		sleepDelay := time.Millisecond * 5
		key := "https://www.test.com"
		value := []byte("this is a test")

		cache := NewCache(interval)
		cache.Add(key, value)

		if _, ok := cache.Get(key); !ok {
			t.Errorf("expected cache entry with key: %s", key)
			return
		}

		time.Sleep(interval + sleepDelay)

		if _, ok := cache.Get(key); ok {
			t.Errorf("expected cache entry with key: %s to have been cleared", key)
			return
		}
	})
}
