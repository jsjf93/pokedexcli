package commands

type Config struct {
	Next     string
	Previous string
}

func NewConfig() Config {
	return Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "https://pokeapi.co/api/v2/location-area",
	}
}
