package inmemorycache

// Cashe
type Cache struct {
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
	delete(c.items, key)
}
