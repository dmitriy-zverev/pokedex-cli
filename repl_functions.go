package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for key, value := range supportedFunctions {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}
