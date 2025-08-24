package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dmitriy-zverev/pokedex-cli/cliHandler"
	"github.com/dmitriy-zverev/pokedex-cli/pokecache"
)

func main() {
	cliHandler.InitSupportedFunctions()

	userInput := bufio.NewScanner(os.Stdin)
	userConfig := cliHandler.Config{}
	userCache := pokecache.NewCache(60 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() && len(userInput.Text()) > 0 {
			cleanedInput := cliHandler.CleanInput(userInput.Text())
			if len(cleanedInput) < 1 {
				continue
			}

			userCommand := cleanedInput[0]
			if cmd, ok := cliHandler.SupportedFunctions[userCommand]; ok {
				err := cmd.Callback(&userConfig, &userCache)
				if err != nil {
					fmt.Printf("Error running the command: %v\n", err)
				}
			}
		}
	}
}
