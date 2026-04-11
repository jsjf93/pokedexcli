package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jsjf93/pokedexcli/internal/commands/apiresponses"
)

func (c *Client) ListLocationAreas(url string) (apiresponses.LocationAreasResponse, error) {
	var response apiresponses.LocationAreasResponse

	if cacheResult, ok := c.Cache.Get(url); ok {
		if err := json.Unmarshal(cacheResult, &response); err != nil {
			return apiresponses.LocationAreasResponse{}, err
		}

		return response, nil
	}

	res, err := c.HttpClient.Get(url)

	if res.StatusCode > 200 {
		return apiresponses.LocationAreasResponse{}, fmt.Errorf("network request failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		return apiresponses.LocationAreasResponse{}, fmt.Errorf("error retrieving location areas: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return apiresponses.LocationAreasResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return apiresponses.LocationAreasResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	if err := c.Cache.Add(url, body); err != nil {
		return apiresponses.LocationAreasResponse{}, fmt.Errorf("error caching response: %w", err)
	}

	return response, nil
}
