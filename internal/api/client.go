package api

import (
	"time"

	"github.com/polaski0/pokedexcli/internal/cache"
)

type Client struct {
	cache cache.Cache
    Pokedex map[string]Pokemon
}

func NewClient(d time.Duration) Client {
	return Client{
        cache: cache.NewCache(d),
        Pokedex: make(map[string]Pokemon),
    }
}
