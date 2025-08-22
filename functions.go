package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var finalStrings []string

	splittedText := strings.Split(strings.TrimSpace(text), " ")

	for _, word := range splittedText {
		finalWord := strings.TrimSpace(word)

		if finalWord != "" {
			finalStrings = append(finalStrings, finalWord)
		}
	}

	fmt.Println(finalStrings)

	return finalStrings
}
