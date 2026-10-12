package cache_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/kilia/pipiline-example/cache"
	"github.com/kilia/pipiline-example/settings"
)

func newCache[T any](shards int) *cache.InMemoryCache[T] {
	return cache.NewInMemoryCache[T](settings.Settings{Shards: shards})
}

func TestSetGet(t *testing.T) {
	c := newCache[string](4)
	c.Set("one", "route", time.Minute)

	got, ok := c.Get("one")
	if !ok || got != "route" {
		t.Fatalf("got %q %v", got, ok)
	}
}

func TestGetMiss(t *testing.T) {
	c := newCache[string](4)

	got, ok := c.Get("missing")
	if ok || got != "" {
		t.Fatalf("got %q %v", got, ok)
	}
}

func TestZeroValueIsHit(t *testing.T) {
	c := newCache[int](4)
	c.Set("n", 0, time.Minute)

	got, ok := c.Get("n")
	if !ok || got != 0 {
		t.Fatalf("got %d %v", got, ok)
	}
}

func TestOverwrite(t *testing.T) {
	c := newCache[string](4)
	c.Set("k", "a", time.Minute)
	c.Set("k", "b", time.Minute)

	got, ok := c.Get("k")
	if !ok || got != "b" {
		t.Fatalf("got %q %v", got, ok)
	}
}

func TestRemove(t *testing.T) {
	c := newCache[string](4)
	c.Set("k", "v", time.Minute)
	c.Remove("k")

	if _, ok := c.Get("k"); ok {
		t.Fatal("key still present")
	}

	c.Remove("k")
}

func TestTTL(t *testing.T) {
	c := newCache[string](4)
	c.Set("k", "v", 20*time.Millisecond)

	if _, ok := c.Get("k"); !ok {
		t.Fatal("key expired too early")
	}

	time.Sleep(40 * time.Millisecond)

	if _, ok := c.Get("k"); ok {
		t.Fatal("key still present after ttl")
	}
}

func TestKeysRoundTripAcrossShards(t *testing.T) {
	c := newCache[string](11)
	for i := range 100 {
		k := fmt.Sprintf("k-%d", i)
		c.Set(k, k, time.Minute)
	}
	for i := range 100 {
		k := fmt.Sprintf("k-%d", i)
		got, ok := c.Get(k)
		if !ok || got != k {
			t.Fatalf("%s: got %q %v", k, got, ok)
		}
	}
}

func TestNewPanicsWhenShardsBelowOne(t *testing.T) {
	for _, n := range []int{0, -1} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic")
				}
			}()
			newCache[int](n)
		})
	}
}

func TestConcurrentSetGet(t *testing.T) {
	c := newCache[int](8)
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k := strconv.Itoa(i)
			c.Set(k, i, time.Minute)
			got, ok := c.Get(k)
			if !ok || got != i {
				t.Errorf("key %s: got %d %v", k, got, ok)
			}
		}()
	}
	wg.Wait()
}

func TestSetDuringExpiredGetKeepsNewValue(t *testing.T) {
	c := newCache[int](1)
	const key = "k"

	for i := range 100 {
		c.Set(key, 1, time.Nanosecond)
		time.Sleep(time.Millisecond)

		done := make(chan struct{})
		go func() {
			defer close(done)
			c.Get(key)
		}()
		c.Set(key, 2, time.Hour)
		<-done

		got, ok := c.Get(key)
		if !ok || got != 2 {
			t.Fatalf("iteration %d: got %d %v", i, got, ok)
		}
		c.Remove(key)
	}
}