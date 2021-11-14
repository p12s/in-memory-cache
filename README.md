![License](https://img.shields.io/github/license/p12s/in-memory-cache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/in-memory-cache?style=plastic)
[![Codecov](https://codecov.io/gh/p12s/in-memory-cache/branch/master/graph/badge.svg?token=0VP8CWJB7A)](https://codecov.io/gh/p12s/in-memory-cache)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/in-memory-cache)](https://goreportcard.com/report/github.com/p12s/in-memory-cache)
<img src="https://github.com/p12s/in-memory-cache/workflows/lint-build-test/badge.svg?branch=master">

# In-memory cache
Task description is [here](task.md)

## Performance and implementation
This cache implementation is based on the sync.Map.  
The solutions based on a hash table promise operations **add**, **get** and **remove** performed in **constant** (O[1]) time.  
For this library, I initially took sync.Mutex + map[string]interface{}, however it showed problems with concurrent deletion of elements and poor performance when deleting itself.  
At the same time sync.Map performed well in these cases (the library, which is taken for performance comparison, was based on sync.Map).   
As a result, I choice sync.Map and I ended up with exactly the same library as the one taken as an example.  
Loser.  

## Example
Installation: 
```go
go get -u github.com/p12s/in-memory-cache
```
Using:
```go
package main

import (
	"fmt"
	"time"
	cache "github.com/p12s/in-memory-cache"
)

func main() {
	cleanExpiredPeriod := 1 * time.Second
	cache := cache.New(cleanExpiredPeriod)
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

| Package                                  	| Method            	| quantity  	| ns/op 	| B/op 	| allocs/op 	| fast % 	|
|------------------------------------------	|-------------------	|-----------	|-------	|------	|-----------	|--------	|
| [Cache](https://github.com/akyoto/cache) 	|                   	|           	|       	|      	|           	|        	|
|                                          	| New               	| 6379844   	| 166.9 	| 368  	| 6         	|        	|
|                                          	| Set/SetWithExpire 	| 6264136   	| 192.4 	| 40   	| 2         	|        	|
|                                          	| Get               	| 420062818 	| 2.911 	| 0    	| 0         	|        	|
|                                          	| Delete            	| 855741927 	| 1.385 	| 0    	| 0         	|        	|
| This package                             	|                   	|           	|       	|      	|           	|        	|
|                                          	| New               	| 6436862   	| 175.6 	| 368  	| 6         	| ~      	|
|                                          	| Set/SetWithExpire 	| 5521599   	| 199.7 	| 56   	| 3         	| ~      	|
|                                          	| Get               	| 424934931 	| 2.824 	| 0    	| 0         	| ~      	|
|                                          	| Delete            	| 842618384 	| 1.417 	| 0    	| 0         	| ~      	|
