package geecache

import (
	"geecache/geecache/lru"
	"sync"
)

type cache struct {
	mu sync.Mutex
	lru *lru.Cache
	cacheBytes int64
}

// Add 添加缓存
func (c *cache) Add(key string, value ByteView)  {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	c.lru.Add(key, value)
}

// Get 获取缓存
func (c *cache) Get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		return
	}

	v, ok := c.lru.Get(key)
	if ok {
		return v.(ByteView), ok
	}

	return
}



