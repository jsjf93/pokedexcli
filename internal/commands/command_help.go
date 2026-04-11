package commands

import (
	"fmt"
)

func CommandHelp(*Config, string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for _, command := range NewRegistry() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	return nil
}
