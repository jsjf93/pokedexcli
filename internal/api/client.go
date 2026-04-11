package api

import (
	"net/http"
	"time"

	"github.com/jsjf93/pokedexcli/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	Cache      pokecache.SafeCache
}

func NewClient(interval time.Duration) Client {
	return Client{
		HttpClient: http.Client{},
		Cache:      pokecache.NewCache(interval),
	}
}
