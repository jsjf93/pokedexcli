package commands

import (
	"strings"

	"github.com/jsjf93/pokedexcli/internal/api"
)

const BaseUrl string = "https://pokeapi.co/api/v2/location-area"

type Config struct {
	Client   api.Client
	Next     string
	Previous string
}

func NewConfig(client api.Client) Config {
	return Config{
		Client:   client,
		Next:     BaseUrl,
		Previous: BaseUrl,
	}
}

func (c *Config) UpdateUrls(next, previous string) {
	nextQuery := strings.Split(next, "?")
	previousQuery := strings.Split(previous, "?")

	if len(nextQuery) != 2 {
		c.Next = BaseUrl
	} else {
		c.Next = next
	}

	if len(previousQuery) != 2 {
		c.Previous = BaseUrl
	} else {
		c.Previous = previous
	}
}
