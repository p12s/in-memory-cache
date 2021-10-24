package inmemorycache

import (
	"sync"
	"time"
)

// item - cache item
type item struct {
	data    interface{}
	expires int64
}

// Cache
type Cache struct {
	close chan struct{}
	items sync.Map
}

// New - constructor
func New(cleaningInterval time.Duration) *Cache {
	cache := Cache{
		close: make(chan struct{}),
	}
	go cache.cleanExpired(cleaningInterval)
	return &cache
}

// cleanExpired - removing expired cache cleaning loop
func (c *Cache) cleanExpired(cleaningInterval time.Duration) {
	ticker := time.NewTicker(cleaningInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now().Unix()
			c.items.Range(func(key, value interface{}) bool {
				item := value.(item) //nolint
				if item.expires != 0 && item.expires < now {
					c.items.Delete(key)
				}
				return true
			})
		case <-c.close:
			return
		}
	}
}

// Set - add value to cache
func (c *Cache) Set(key string, value interface{}) {
	c.items.Store(key, item{
		data: value,
	})
}

// SetWithExpire - add value to cache with expiting time
func (c *Cache) SetWithExpire(key string, value interface{}, ttl time.Duration) {
	c.items.Store(key, item{
		data:    value,
		expires: time.Now().Add(ttl).Unix(),
	})
}

// Get - get value from cache
func (c *Cache) Get(key string) interface{} {
	value, exists := c.items.Load(key)
	if !exists {
		return nil
	}

	item := value.(item) //nolint
	if item.expires != 0 && item.expires < time.Now().Unix() {
		return nil
	}

	return item.data
}

// Delete - remove value from cache
func (c *Cache) Delete(key string) {
	c.items.Delete(key)
}

// Close - removing expired cache cleaning loop and all cache items
func (c *Cache) Close() {
	c.close <- struct{}{}
	c.items = sync.Map{}
}
