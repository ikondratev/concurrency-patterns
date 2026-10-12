package cache

import (
	"hash/maphash"
	"sync"
	"time"

	"github.com/kilia/pipiline-example/settings"
)

type InMemoryCache[T any] struct {
	shurds []shard[T]
	seed   maphash.Seed
}

type shard[T any] struct {
	data 	map[string]data[T]
	mu 	 	sync.RWMutex
	next 	uint64
}

type data[T any] struct {
	value 	  T
	expiredAt time.Time
	version   uint64
}

func NewInMemoryCache[T any](s settings.Settings) *InMemoryCache[T] {
	shards := s.Shards
	if shards < 1 {
		panic("shard below 1")
	}

	shurds := make([]shard[T], 0, shards)
	for range shards {
		shurds = append(shurds, shard[T]{data: make(map[string]data[T])},)
	}

	return &InMemoryCache[T]{
		shurds: shurds,
		seed: 	maphash.MakeSeed(),
	}
}

func (c *InMemoryCache[T]) Set(k string, v T, ttl time.Duration) {
	c.shurds[c.index(k)].set(k, v, ttl)
}

func (c *InMemoryCache[T]) Get(k string) (T, bool) {
	return c.shurds[c.index(k)].get(k)
}

func (c *InMemoryCache[T]) Remove(k string) {
	c.shurds[c.index(k)].remove(k)
}

func (s *shard[T]) set(k string, v T, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.next++
	s.data[k] = data[T]{
		value: 	   v,
		expiredAt: time.Now().Add(ttl),
		version:   s.next,
	}
}

func (s *shard[T]) get(k string) (T, bool) {
	s.mu.RLock()
	v, ok := s.data[k]
	s.mu.RUnlock()

	if !ok {
		var zero T
		return zero, false
	}

	if time.Now().After(v.expiredAt) {
		s.checkForDelete(k, v.version)
		var zero T
		return zero, false
	}

	return v.value, ok
}

func (s *shard[T]) remove(k string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, k)
}

func(s *shard[T]) checkForDelete(k string, version uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.data[k]
	if ok && v.version == version {
		delete(s.data, k)
	}
}

func (c *InMemoryCache[T]) index(k string) int {
	return int(maphash.String(c.seed, k) % uint64(len(c.shurds)))
}
