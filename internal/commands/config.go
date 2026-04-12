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
	Collection map[string]Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		Collection: make(map[string]Pokemon),
	}
}

type Pokemon struct {
	Name string
}

func (p *Pokedex) Add(key string, pokemon Pokemon) {
	if _, found := p.Collection[key]; !found {
		p.Collection[key] = pokemon
	}
}

func (p *Pokedex) Get(key string) (*Pokemon, bool) {
	if pokemon, found := p.Collection[key]; !found {
		return &pokemon, true
	}

	return nil, false
}
