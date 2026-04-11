package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jsjf93/pokedexcli/internal/api"
	"github.com/jsjf93/pokedexcli/internal/commands"
	"github.com/jsjf93/pokedexcli/internal/repl"
)

func main() {
	commandRegistry := commands.NewRegistry()
	interval := time.Duration(60 * time.Second)
	client := api.NewClient(interval)
	config := commands.NewConfig(client)
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
		arg := ""

		if len(words) > 1 {
			arg = words[1]
		}

		command, found := commandRegistry[commandKey]

		if !found {
			fmt.Println("Unknown command")
		} else {
			err := command.Callback(&config, arg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
