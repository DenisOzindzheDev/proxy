package testing

import (
	"fmt"
	"rproxy/pkg/cache"
)

func testCacheChannels() {
	cache := cache.NewLru(15)
	var s = make(chan interface{})
	go func() {
		cache.SetKey("12", "12")
		s <- cache.GetKey("12")
	}()
	val := <-s
	fmt.Println(val)
}
