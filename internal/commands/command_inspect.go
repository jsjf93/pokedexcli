package commands

import (
	"errors"
	"strings"
)

func CommandInspect(config *Config, pokemon string) error {
	searchTerm := strings.ToLower(strings.TrimSpace(pokemon))
	if searchTerm == "" {
		return errors.New("pokemon name is required")
	}

	match, found := config.Pokedex.Get(searchTerm)
	if !found {
		return errors.New("you have not caught that pokemon")
	}

	match.Print()

	return nil
}
