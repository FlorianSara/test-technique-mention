package main

import (
	"fmt"
	"test2/pkg/customCache"
)

func main() {
	var myCache customCache.Cache

	myCache = customCache.NewMemCache()
	// myCache = customCache.NewFileCache("/tmp/customCache")

	myCache.Set("test1", "test")
	myCache.Set("test2", 32)
	myCache.Set("test3", []string{"first", "second"})

	fmt.Printf("result: %v\n", myCache.Get("test1"))
	fmt.Printf("result: %v\n", myCache.Get("test2"))
	fmt.Printf("result: %v\n", myCache.Get("test3"))
	fmt.Printf("result: %v\n", myCache.Get("test4"))
}
