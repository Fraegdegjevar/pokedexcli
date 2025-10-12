package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	// Initialise a cache without using NewCache so we don't
	// have deletion of 'old' entries
	cases := []struct {
		input    map[string][]byte
		expected []byte
	}{
		{
			input:    map[string][]byte{"first": []byte("hello")},
			expected: []byte{104, 101, 108, 108, 111},
		},
		{
			input:    map[string][]byte{"second": {}},
			expected: []byte{},
		},
		{
			input:    map[string][]byte{"third": []byte(" ")},
			expected: []byte{32},
		},
		{
			input:    map[string][]byte{"fourth": []byte("543!@")},
			expected: []byte{53, 52, 51, 33, 64},
		},
	}

	for _, c := range cases {
		cache := &Cache{
			Entries: make(map[string]cacheEntry),
		}
		for key, value := range c.input {
			cache.Add(key, value)
			if !bytes.Equal(cache.Entries[key].val, c.expected) {
				t.Errorf("Expected cache entry byte val: %v, Actual entry: %v ", c.expected, cache.Entries[key].val)
			}
		}
	}
}

func TestCacheGet(t *testing.T) {

	cases := []struct {
		input         *Cache
		key           string
		expectedVal   []byte
		expectedExist bool
	}{
		{
			input: &Cache{
				Entries: map[string]cacheEntry{"test": {createdAt: time.Now(), val: []byte("test")}},
			},
			key:           "test",
			expectedVal:   []byte("test"),
			expectedExist: true,
		},
		{
			input: &Cache{
				Entries: map[string]cacheEntry{"1": {createdAt: time.Now(), val: []byte("1")}},
			},
			key:           "",
			expectedVal:   []byte{},
			expectedExist: false,
		},
	}
	for _, c := range cases {
		val, exist := c.input.Get(c.key)
		if !bytes.Equal(c.expectedVal, val) {
			t.Errorf("Expected to Get value %v, got %v", c.expectedVal, val)
		}

		if exist != c.expectedExist {
			t.Errorf("Expected existence of key to be %v, but got: %v", c.expectedExist, exist)
		}
	}
}

func TestCacheReapLoop(t *testing.T) {
	cache := NewCache(1 * time.Millisecond)

	cache.Add("one", []byte{32})
	//Allow reapLoop to tick even if delayed slightly
	time.Sleep(3 * time.Millisecond)

	if _, exist := cache.Get("one"); exist {
		t.Errorf("Expected cache entry one to be removed by reapLoop but Get() returned %v", exist)
	}

	cache.Add("two", []byte{32})
	if _, exist := cache.Get("two"); !exist {
		t.Errorf("Expected cache entry two to exist as not old enough to be reaped but Get() failed to find key")
	}

}
