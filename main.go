package main

import (
	"fmt"
	"github.com/krls256/knowlage-base-editor/internal/cache"
)

func main()  {
	cache.SecretFileCache.Set("hello", "World")
	fmt.Println(cache.SecretFileCache.Get("hello"))
}
