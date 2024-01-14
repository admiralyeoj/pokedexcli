package pokeCache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mu:           &sync.Mutex{},
		cache:        make(map[string]cacheEntry),
		reapInterval: interval,
	}

	go c.reapLoop()

	return c
}
