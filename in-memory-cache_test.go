package inmemorycache

import (
	"testing"
	"time"

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
			assert.Equal(t, nil, cache.Get(tt.key))

			// now key-value should exists
			cache.Set(tt.key, tt.value)
			assert.Equal(t, tt.value, cache.Get(tt.key))

			// againt, key-value shouldn't exists
			cache.Delete(tt.key)
			assert.Equal(t, nil, cache.Get(tt.key))
		})
	}
}

func TestSetWithExpire(t *testing.T) {
	t.Parallel()

	cache := New()

	tests := []struct {
		name, key string
		value     interface{}
		existsSec int
	}{
		{"Can set the negative expiration time to -5 sec - it will be taken as 0 sec", "int", 42, -5},
		{"Can set the expiration time to 1 sec", "int1", 42, 1},
		{"Can set the expiration time to 2 sec", "int2", 42, 2},
		{"Can set the expiration time to 3 sec", "int3", 42, 3},
		{"Can set the expiration time to 4 sec", "int4", 42, 4},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// key-value shouldn't exists
			assert.Equal(t, nil, cache.Get(tt.key))

			// now key-value should exists (if ttl time > 0)
			cache.SetWithExpire(tt.key, tt.value, time.Second*time.Duration(tt.existsSec))
			if tt.existsSec > 0 {
				assert.Equal(t, tt.value, cache.Get(tt.key))
			} else {
				assert.Equal(t, nil, cache.Get(tt.key))
			}

			// key-value still should exists
			if tt.existsSec >= 0 {
				for i := 0; i < tt.existsSec; i++ {
					time.Sleep(time.Second * 1)
					assert.Equal(t, tt.value, cache.Get(tt.key))
				}
			}

			// time variation
			time.Sleep(time.Second * 1)

			// key-value shouldn't exists
			assert.Equal(t, nil, cache.Get(tt.key))
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

func BenchmarkSetWithExpire(b *testing.B) {
	cache := New()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.SetWithExpire("Hello", "World", time.Second*1)
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
