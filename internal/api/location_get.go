package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
)

const baseUrl string = "https://pokeapi.co/api/v2/location-area/"

func (c *Client) GetLocation(locationArea string) (apiresponses.LocationAreaResponse, error) {
	url := baseUrl + locationArea
	var response apiresponses.LocationAreaResponse

	if cacheEntry, ok := c.Cache.Get(url); ok {
		if err := json.Unmarshal(cacheEntry, &response); err != nil {
			return apiresponses.LocationAreaResponse{}, err
		}
		return response, nil
	}

	res, err := c.HttpClient.Get(url)

	if res.StatusCode > 200 {
		return apiresponses.LocationAreaResponse{}, fmt.Errorf("network request failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		return apiresponses.LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return apiresponses.LocationAreaResponse{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return apiresponses.LocationAreaResponse{}, err
	}

	if err := c.Cache.Add(url, body); err != nil {
		return apiresponses.LocationAreaResponse{}, err
	}

	return response, nil
}
