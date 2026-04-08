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
	}

	return registry
}
