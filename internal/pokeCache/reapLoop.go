package pokeCache

import (
	"fmt"
	"time"
)

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if time.Since(entry.createdAt) > c.reapInterval {
					fmt.Println("Deleting " + key)
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
