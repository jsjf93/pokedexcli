package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jsjf93/pokedexcli/internal/commands"
	"github.com/jsjf93/pokedexcli/internal/repl"
)

func main() {
	commandRegistry := commands.NewRegistry()
	config := commands.NewConfig()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := repl.CleanInput(scanner.Text())

		if len(words) == 0 {
			fmt.Println("Input is required")
			continue
		}

		commandKey := words[0]

		command, found := commandRegistry[commandKey]

		if !found {
			fmt.Println("Unknown command")
		} else {
			err := command.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
