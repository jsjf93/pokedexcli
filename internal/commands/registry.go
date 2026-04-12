package commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, arg string) error
}

type Registry = map[string]cliCommand

func NewRegistry() Registry {
	registry := Registry{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Retrieves the next 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Retrieves the previous 20 location areas in the Pokemon world",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Retrieves the Pokemon you can encounter in a given area. Example: explore canalave-city-area",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch a Pokemon. Example: catch pikachu",
			Callback:    CommandCatch,
		},
	}

	return registry
}
