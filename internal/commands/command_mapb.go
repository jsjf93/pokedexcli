package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
)

func CommandMapb(config *Config) error {
	url := config.Previous

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error retrieving location areas: %w", err)
	}

	defer res.Body.Close()

	var locationAreaResponse apiresponses.LocationAreaResponse

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreaResponse); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}

	config.Next = locationAreaResponse.Next
	config.Previous = locationAreaResponse.Previous

	return nil
}
