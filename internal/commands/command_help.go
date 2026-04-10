package commands

import (
	"fmt"

	"github.com/jsjf93/pokedexcli/internal/pokecache"
)

func CommandHelp(*Config, pokecache.SafeCache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for _, command := range NewRegistry() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	return nil
}
