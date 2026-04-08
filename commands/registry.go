package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type Registry = map[string]cliCommand

func NewRegistry() map[string]cliCommand {
	registry := map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
	}

	return registry
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for _, command := range NewRegistry() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	return nil
}
