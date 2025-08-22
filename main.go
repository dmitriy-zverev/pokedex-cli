package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() && len(userInput.Text()) > 0 {
			userString := cleanInput(userInput.Text())[0]
			fmt.Printf("Your command was: %s\n", userString)
		}
	}
}
