package inmemorycache

import "sync"

// Cashe
type Cache struct {
	mu    sync.Mutex
	items map[string]interface{}
}

// Constructor
func New() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

// Set - add value to cache
func (c *Cache) Set(key string, item interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = item
}

// Get - get value from cache
func (c *Cache) Get(key string) interface{} {
	if item, ok := c.items[key]; ok {
		return item
	}
	return nil
}

// Delete - remove value from cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
