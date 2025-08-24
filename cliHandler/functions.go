package cliHandler

import (
	"strings"
)

func CleanInput(text string) []string {
	var finalStrings []string

	splittedText := strings.Split(strings.TrimSpace(text), " ")

	for _, word := range splittedText {
		finalWord := strings.ToLower(strings.TrimSpace(word))

		if finalWord != "" {
			finalStrings = append(finalStrings, finalWord)
		}
	}

	return finalStrings
}
