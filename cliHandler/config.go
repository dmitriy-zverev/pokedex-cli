package cliHandler

import "errors"

type Config struct {
	Pokedex     map[string]Pokemon
	Command     int
	FullCommand []string
	Next        string
	Previous    string
}

func updateConfig(command int, config *Config) error {
	if command == config.Command {
		return nil
	}

	switch command {
	case REPL_HELP, REPL_EXIT:
		config.Next = ""
		config.Previous = ""
		config.Command = command
	case REPL_MAP:
		config.Previous = POKEAPI_BASE_URL + "/location-area/0"
		config.Next = POKEAPI_BASE_URL + "/location-area/1"
		config.Command = command
	default:
		return errors.New("error: unknown command")
	}

	return nil
}
