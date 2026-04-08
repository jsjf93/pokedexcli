package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jsjf93/pokedexcli/commands"
	"github.com/jsjf93/pokedexcli/repl"
)

func main() {
	commands := commands.NewRegistry()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := repl.CleanInut(scanner.Text())

		if len(words) == 0 {
			fmt.Println("Input is required")
			continue
		}

		commandKey := words[0]

		command, found := commands[commandKey]

		if !found {
			fmt.Println("Unknown command")
		} else {
			command.Callback()
		}
	}
}
