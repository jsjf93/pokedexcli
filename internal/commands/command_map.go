package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CommandMap(config *Config) error {
	url := config.Next

	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error retrieving location areas: %w", err)
	}

	defer res.Body.Close()

	var locationAreaResponse locationAreaResponse
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

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
