package commands

import (
	"strings"

	"github.com/jsjf93/pokedexcli/internal/api"
	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
)

const BaseUrl string = "https://pokeapi.co/api/v2/location-area"

type Config struct {
	Client   api.Client
	Next     string
	Previous string
	Pokedex  Pokedex
}

func NewConfig(client api.Client) Config {
	return Config{
		Client:   client,
		Next:     BaseUrl,
		Previous: BaseUrl,
		Pokedex:  NewPokedex(),
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

type Pokedex struct {
	collection map[string]apiresponses.PokemonResponse
}

func NewPokedex() Pokedex {
	return Pokedex{
		collection: make(map[string]apiresponses.PokemonResponse),
	}
}

func (p *Pokedex) Add(key string, pokemon apiresponses.PokemonResponse) {
	if _, found := p.collection[key]; !found {
		p.collection[key] = pokemon
	}
}

func (p *Pokedex) Get(key string) (apiresponses.PokemonResponse, bool) {
	pokemon, found := p.collection[key]
	return pokemon, found
}
