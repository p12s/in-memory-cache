![License](https://img.shields.io/github/license/p12s/in-memory-cache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/in-memory-cache?style=plastic)
[![Codecov](https://codecov.io/gh/p12s/in-memory-cache/branch/master/graph/badge.svg?token=0VP8CWJB7A)](https://codecov.io/gh/p12s/in-memory-cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/in-memory-cache)](https://goreportcard.com/report/github.com/p12s/in-memory-cache)
<img src="https://github.com/p12s/in-memory-cache/workflows/lint-build/badge.svg?branch=master">

# In-memory cache
Task description is [here](task.md)

## Performance and implementation
This cache implementation is based on the standard go structure - map[string]interface{}.  
In turn, the map type is based on a hash table, so we consider that **adding**, **getting** and **removing** values is performed in **constant time** (O[1]).  
However, when the number of elements equals the capacity of the dictionary, the entire structure will be copied to a new location in memory,  
with a 2-fold increase in capacity (on 64-bit machines, when the number reaches 2048, it will be 1.25 times).  
Therefore, this self-written cache is not optimal, I would not recommend using it in production.

## Example
Installation: 
```
go get -u https://github.com/p12s/in-memory-cache
```
Использование:
```
cache := cache.New()

userId := cache.Get("userId")   // if the key doesn't exist, returns "nil"
fmt.Println(userId)

cache.Set("userId", 42)         // if the key already exists, it will overwrite
userId := cache.Get("userId")
fmt.Println(userId)

cache.Delete("userId")
userId := cache.Get("userId")
fmt.Println(userId)
```


Maps are backed by hash tables.
Add, get and delete operations run in constant expected time. The time complexity for the add operation is amortized.
The comparison operators == and != must be defined for the key type.
