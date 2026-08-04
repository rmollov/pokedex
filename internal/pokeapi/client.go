package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
	Pokedex    map[string]Pokemon
}

// NewClient -
func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokecache.NewCache(interval),
		Pokedex: make(map[string]Pokemon),
	}
}
