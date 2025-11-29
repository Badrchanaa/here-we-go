package main

import (
	"fmt"

	"github.com/Badrchanaa/here-we-go/system-design/internal/lru_cache"
)

func main() {
	cache := lru_cache.Constructor(2)
	cache.Put(4, 10)
	cache.Put(2, 50)
	cache.Put(4, 30)
	val := cache.Get(4)
	fmt.Println("cache[4] = ", val)
}
