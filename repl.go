package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	stringLower := strings.ToLower(text)
	return strings.Fields(stringLower)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//Notice lack of newline
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}
		input := scanner.Text()
		if input == "" {
			continue
		}
		firstWord := cleanInput(input)[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
