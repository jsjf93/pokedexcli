package commands

import (
	"fmt"
	"os"

	"github.com/jsjf93/pokedexcli/internal/pokecache"
)

func CommandExit(*Config, pokecache.SafeCache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
