package pokeCache

import "time"

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
