package commands

import (
	"fmt"
)

func CommandMap(config *Config) error {
	err := getLocationArea(config, config.Next)
	if err != nil {
		return err
	}

	return nil
}

func CommandMapb(config *Config) error {
	err := getLocationArea(config, config.Previous)
	if err != nil {
		return err
	}

	return nil
}

func getLocationArea(
	config *Config,
	url string,
) error {
	res, err := config.Client.ListLocationAreas(url)
	if err != nil {
		return err
	}

	for _, locationArea := range res.Results {
		fmt.Println(locationArea.Name)
	}

	config.UpdateUrls(res.Next, res.Previous)

	return nil
}
