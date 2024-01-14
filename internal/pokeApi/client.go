package pokeApi

import (
	"net/http"
	"time"

	"github.com/admiralyeoj/pokedexcli/internal/pokeCache"
)

// Client -
type Client struct {
	cache      pokeCache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokeCache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
