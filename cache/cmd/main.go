package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kilia/pipiline-example/cache"
	"github.com/kilia/pipiline-example/settings"
)

func main() {
	set := settings.Settings{Shards: 11}
	c := cache.NewInMemoryCache[string](set)
	c.Set("one", "route", 5 * time.Minute)
	v, _ := c.Get("one")
	fmt.Println("Result:", v)

	var wg sync.WaitGroup
	wg.Add(11)

	for i := range 11 {
		go func() {
			defer wg.Done(

			)
			k := fmt.Sprintf("Value %d", i + 1)
			c.Set(k, k, 5 * time.Second)
		}()
	}

	wg.Wait()

	fmt.Println("Vp:", c)
}