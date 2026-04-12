package commands

import "fmt"

func CommandPokedex(config *Config, _ string) error {
	if len(config.Pokedex.collection) == 0 {
		return fmt.Errorf("your Pokedex is empty!")
	}

	fmt.Println("Your Pokedex")

	for _, value := range config.Pokedex.collection {
		fmt.Printf(" - %s\n", value.Name)
	}

	return nil
}
