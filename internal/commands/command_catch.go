package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func CommandCatch(config *Config, pokemon string) error {
	searchTerm := strings.ToLower(strings.TrimSpace(pokemon))
	if searchTerm == "" {
		return errors.New("pokemon name is required")
	}

	res, err := config.Client.GetPokemon(searchTerm)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	isSuccess := CatchAttempt(res.BaseExperience)

	if isSuccess {
		fmt.Printf("%s was caught!\n", res.Name)
		config.Pokedex.Add(strings.ToLower(res.Name), res)
	} else {
		fmt.Printf("%s escaped!\n", res.Name)
	}

	return nil
}

func CatchAttempt(baseExperience int) bool {
	if baseExperience <= 0 {
		return true
	}
	randomCatchFactor := rand.Intn(baseExperience * 2)

	return randomCatchFactor >= baseExperience
}
