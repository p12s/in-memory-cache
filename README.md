# In-memory cache
Полное описание задачи [тут](task.md)

## Выполнение
Установка пакета: 
```
go get -u https://github.com/p12s/in-memory-cache
```
Использование:
```
cache := cache.New()

cache.Set("userId", 42)
userId := cache.Get("userId")

fmt.Println(userId)

cache.Delete("userId")
userId := cache.Get("userId")

fmt.Println(userId)
```
