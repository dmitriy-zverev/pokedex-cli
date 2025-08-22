package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	initSupportedFunctions()
	userInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() && len(userInput.Text()) > 0 {
			cleanedInput := cleanInput(userInput.Text())
			if len(cleanedInput) < 1 {
				continue
			}

			userCommand := cleanedInput[0]
			if cmd, ok := supportedFunctions[userCommand]; ok {
				cmd.callback()
			}
		}
	}
}
