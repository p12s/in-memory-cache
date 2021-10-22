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

// Cashe
type Cache struct {
	mu    sync.Mutex // TODO need to experiment with: sync.Map RWMutex
	items map[string]item
}

// Constructor
func New() *Cache {
	return &Cache{
		items: make(map[string]item),
	}
}

// Set - add value to cache
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.items[key] = item{
		data: value,
	}
	c.mu.Unlock() // without defer, because it's add overhead ~200 ns
}

// SetWithExpire - add value to cache with expiting time
func (c *Cache) SetWithExpire(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.items[key] = item{
		data:    value,
		expires: time.Now().Add(ttl).Unix(),
	}
	c.mu.Unlock()
}

// Get - get value from cache
func (c *Cache) Get(key string) interface{} {
	if _, ok := c.items[key]; !ok {
		return nil
	}

	if c.items[key].expires != 0 && c.items[key].expires < time.Now().Unix() {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil
	}

	return c.items[key].data
}

// Delete - remove value from cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}
