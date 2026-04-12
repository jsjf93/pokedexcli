# pokedexcli

I'm trying to learn a bit of Go and went through this guided project from [BootDev](https://www.boot.dev)
`pokedexcli` is a small Go command-line application that uses the [PokeAPI](https://pokeapi.co) to explore Pokemon world data, catch Pokemon, and keep a local in-memory Pokedex during a session.

## Features

- Browse Pokemon location areas with forward and backward paging
- Explore a location area to see which Pokemon can be encountered there
- Attempt to catch Pokemon by name
- Inspect Pokemon you have already caught
- List your caught Pokemon in a session Pokedex

## Requirements

- Go `1.26.1`

## Run

From the project root:

```bash
go run .
```

You should see the prompt:

```text
Pokedex >
```

## Commands

- `help` - Show the available commands
- `map` - Show the next 20 location areas
- `mapb` - Show the previous 20 location areas
- `explore <location-area>` - List Pokemon encounters for a location area
- `catch <pokemon>` - Attempt to catch a Pokemon
- `inspect <pokemon>` - Show details for a caught Pokemon
- `pokedex` - List Pokemon you have caught in the current session
- `exit` - Exit the application

## Example Session

```text
Pokedex > help
Pokedex > map
Pokedex > explore canalave-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```

## Notes

- Caught Pokemon are stored in memory only for the current run of the program
- Catch success is based on the Pokemon's base experience value
- Data is fetched from the public PokeAPI service