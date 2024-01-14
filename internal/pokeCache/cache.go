package pokeCache

import (
	"sync"
	"time"
)

type Cache struct {
	mu           *sync.Mutex
	cache        map[string]cacheEntry
	reapInterval time.Duration // Interval for reaping old entries
}
