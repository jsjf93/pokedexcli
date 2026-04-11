package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
)

func (c *Client) GetPokemon(name string) (apiresponses.PokemonResponse, error) {
	url := baseUrl + "pokemon/" + name

	var response apiresponses.PokemonResponse

	if cacheEntry, ok := c.Cache.Get(url); ok {
		if err := json.Unmarshal(cacheEntry, &response); err != nil {
			return apiresponses.PokemonResponse{}, err
		}
		return response, nil
	}

	res, err := c.HttpClient.Get(url)

	if err != nil {
		return apiresponses.PokemonResponse{}, err
	}

	if res.StatusCode > 200 {
		return apiresponses.PokemonResponse{}, fmt.Errorf("network request failed with status code: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return apiresponses.PokemonResponse{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return apiresponses.PokemonResponse{}, err
	}

	if err := c.Cache.Add(url, body); err != nil {
		return apiresponses.PokemonResponse{}, err
	}

	return response, nil
}
