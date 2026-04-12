package apiresponses

import "fmt"

type PokemonResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	IsDefault      bool    `json:"is_default"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

func (p PokemonResponse) Print() {
	stats := make(map[string]int, len(p.Stats))
	for _, stat := range p.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}

	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
  -hp: %d
  -attack: %d
  -defense: %d
  -special-attack: %d
  -special-defense: %d
  -speed: %d
Types:
`,
		p.Name,
		p.Height,
		p.Weight,
		stats["hp"],
		stats["attack"],
		stats["defense"],
		stats["special-attack"],
		stats["special-defense"],
		stats["speed"])

	for _, pokemonType := range p.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
}
