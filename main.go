package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dmitriy-zverev/pokedex-cli/cliHandler"
)

func main() {
	cliHandler.InitSupportedFunctions()

	userInput := bufio.NewScanner(os.Stdin)
	userConfig := cliHandler.Config{}

	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() && len(userInput.Text()) > 0 {
			cleanedInput := cliHandler.CleanInput(userInput.Text())
			if len(cleanedInput) < 1 {
				continue
			}

			userCommand := cleanedInput[0]
			if cmd, ok := cliHandler.SupportedFunctions[userCommand]; ok {
				err := cmd.Callback(&userConfig)
				if err != nil {
					fmt.Printf("Error running the command: %v\n", err)
				}
			}
		}
	}
}
