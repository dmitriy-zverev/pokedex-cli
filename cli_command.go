package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedFunctions map[string]cliCommand

func initSupportedFunctions() {
	supportedFunctions = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
