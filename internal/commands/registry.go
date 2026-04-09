package commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
	}

	return registry
}
