package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jsjf93/pokedexcli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := repl.CleanInut(scanner.Text())

		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		} else {
			fmt.Println("Unable to process your command")
		}
	}
}
