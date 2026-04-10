package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
	"github.com/jsjf93/pokedexcli/internal/pokecache"
)

func CommandMap(config *Config, cache pokecache.SafeCache) error {
	err := getLocationArea(config, &cache, config.Next)
	if err != nil {
		return err
	}

	return nil
}

func CommandMapb(config *Config, cache pokecache.SafeCache) error {
	err := getLocationArea(config, &cache, config.Previous)
	if err != nil {
		return err
	}

	return nil
}

func getLocationArea(
	config *Config,
	cache *pokecache.SafeCache,
	url string,
) error {
	var locationAreaResponse apiresponses.LocationAreaResponse

	if cacheResult, ok := cache.Get(url); ok {
		json.Unmarshal(cacheResult, &locationAreaResponse)
	} else {
		res, err := http.Get(url)

		if res.StatusCode > 200 {
			return fmt.Errorf("network request failed with status code: %d", res.StatusCode)
		}

		if err != nil {
			return fmt.Errorf("error retrieving location areas: %w", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}

		if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}

		if err := cache.Add(url, body); err != nil {
			return fmt.Errorf("error caching response: %w", err)
		}
	}

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}

	config.UpdateUrls(locationAreaResponse.Next, locationAreaResponse.Previous)

	return nil
}
