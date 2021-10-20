package inmemorycache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCache - in-memory cache testing
func TestNew(t *testing.T) {
	t.Parallel()

	cache := New()

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

func BenchmarkNew(b *testing.B) {

	b.ResetTimer()
	b.ReportAllocs()

	var cache *Cache

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache = New() // mem allocs
		}
	})

	_ = cache
}

func BenchmarkGet(b *testing.B) {
	cache := New()
	cache.Set("Hello", "World")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Get("Hello")
		}
	})
}

func BenchmarkSet(b *testing.B) {
	cache := New()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set("Hello", "World")
		}
	})
}

func BenchmarkDelete(b *testing.B) {
	cache := New()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Delete("Hello")
		}
	})
}
