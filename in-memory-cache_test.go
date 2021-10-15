package inmemorycache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCache - in-memory cache testing
func TestNewCache(t *testing.T) {
	cache := NewCache()

	tests := []struct {
		name, key string
		value     interface{}
	}{
		{"Can set an INT to the cache value", "int", 42},
		{"Can set a FLOAT to the cache value", "float", 12.345},
		{"Can set a STRING to the cache value", "string", "hello"},
		{"Can set a BOOL to the cache value", "bool", true},
		{"Can set an ARRAY to the cache value", "array", [3]int{}},
		{"Can set a SLICE to the cache value", "slice", []int{}},
		{"Can set a STRUCT to the cache value", "struct", struct{ name string }{name: "Pol"}},
		{"Can set a POINTER to the cache value", "point", &struct{ name string }{name: "Pol"}},
		{"Can set a MAP to the cache value", "map", map[int]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// key-value shouldn't exists
			value := cache.Get(tt.key)
			assert.Equal(t, value, nil)

			// now key-value should exists
			cache.Set(tt.key, tt.value)
			value = cache.Get(tt.key)
			assert.Equal(t, value, tt.value)

			// againt, key-value shouldn't exists
			cache.Delete(tt.key)
			value = cache.Get(tt.key)

			assert.Equal(t, value, nil)
		})
	}
}
