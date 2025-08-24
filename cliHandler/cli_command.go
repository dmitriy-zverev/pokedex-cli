package cliHandler

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config) error
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
	}
}
