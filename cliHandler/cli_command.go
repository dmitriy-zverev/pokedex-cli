package cliHandler

import (
	"github.com/dmitriy-zverev/pokedex-cli/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, cache *pokecache.Cache) error
}

var SupportedFunctions map[string]CliCommand

func InitSupportedFunctions() {
	SupportedFunctions = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 next locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays 20 previous locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Expores the location area and returns pokemons at this area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catches a pokemon—or not",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspects a pokemon from your pokedex",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists all of the caught pokemons",
			Callback:    commandPokedex,
		},
	}
}
