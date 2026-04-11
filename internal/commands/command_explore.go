package commands

import "fmt"

func CommandExplore(config *Config, locationArea string) error {
	if locationArea == "" {
		return fmt.Errorf("location area is required")
	}

	res, err := config.Client.GetLocation(locationArea)
	if err != nil {
		return err
	}

	for _, pokemon := range res.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
