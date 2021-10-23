![License](https://img.shields.io/github/license/p12s/in-memory-cache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/in-memory-cache?style=plastic)
[![Codecov](https://codecov.io/gh/p12s/in-memory-cache/branch/master/graph/badge.svg?token=0VP8CWJB7A)](https://codecov.io/gh/p12s/in-memory-cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/in-memory-cache)](https://goreportcard.com/report/github.com/p12s/in-memory-cache)
<img src="https://github.com/p12s/in-memory-cache/workflows/lint-build-test/badge.svg?branch=master">

# In-memory cache
Task description is [here](task.md)

## Performance and implementation
This cache implementation is based on the standard go structure - map[string]interface{}.  
In turn, the map type is based on a hash table, so we consider that **adding**, **getting** and **removing** values is performed in **constant time** (O[1]).  
However, when the number of elements equals the capacity of the dictionary, the entire structure will be copied to a new location in memory,  
with a 2-fold increase in capacity (on 64-bit machines, when the number reaches 2048, it will be 1.25 times).  

## Example
Installation: 
```
go get -u github.com/p12s/in-memory-cache
```
Using:
```
package main

import (
	"fmt"
	"time"
	cache "github.com/p12s/in-memory-cache"
)

func main() {
	cache := cache.New()
	fmt.Println(cache.Get("userId"))	// doesn't exist - return "nil"

	cache.Set("userId", 42)
	fmt.Println(cache.Get("userId"))	// 42

	cache.SetWithExpire("number", 43, time.Second * 10)
	fmt.Println(cache.Get("number"))	// 43
	time.Sleep(time.Second * 11)
	fmt.Println(cache.Get("number"))	// nil

	cache.Delete("userId")
	fmt.Println(cache.Get("userId"))	// <nil>
}
```

## Сomparison with other same solution
For comparison, let's take any of the package from [awesome-go](https://github.com/avelino/awesome-go), for example, [In-memory cache](https://github.com/akyoto/cache).  

| Package                                  	| Method        	| Quantity   	| Time         	| Size     	| Allocs      	| %          	|
|------------------------------------------	|---------------	|------------	|--------------	|----------	|-------------	|------------	|
| [Cache](https://github.com/akyoto/cache) 	|               	|            	|              	|          	|             	|            	|
|                                          	| New           	| 6379844    	| 166.9 ns/op  	| 368 B/op 	| 6 allocs/op 	|            	|
|                                          	| Set           	| 6264136    	| 192.4 ns/op  	| 40 B/op  	| 2 allocs/op 	|            	|
|                                          	| Get           	| 420062818  	| 2.911 ns/op  	| 0 B/op   	| 0 allocs/op 	|            	|
|                                          	| Delete        	| 855741927  	| 1.385 ns/op  	| 0 B/op   	| 0 allocs/op 	| ~47 faster 	|
| This package                             	|               	|            	|              	|          	|             	|            	|
|                                          	| New           	| 37835570   	| 29.66 ns/op  	| 80 B/op  	| 2 allocs/op 	| ~5 faster  	|
|                                          	| Set           	| 15231852   	| 77.39 ns/op  	| 0 B/op   	| 0 allocs/op 	| ~2 faster  	|
|                                          	| SetWithExpire 	| 8551774    	| 134.5 ns/op  	| 0 B/op   	| 0 allocs/op 	| ~ same     	|
|                                          	| Get           	| 1000000000 	| 0.4748 ns/op 	| 0 B/op   	| 0 allocs/op 	| ~6 faster  	|
|                                          	| Delete        	| 22668232   	| 52.08 ns/op  	| 0 B/op   	| 0 allocs/op 	| ~47 slower 	|
Testing on 8 cores, i9 2.30GHz.    
Everything, except for items deleting, is no worse than in the compared package.  
