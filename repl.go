package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Fraegdegjevar/pokedexcli/internal/command"
	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
	"github.com/Fraegdegjevar/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{Cache: pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon)}
	supportedCommands := command.GetSupportedCommands()

	for {
		//Notice lack of newline
		fmt.Print("Pokedex > ")

		//Progress scanner to text object or loop again if missing
		if !scanner.Scan() {
			continue
		}
		//If blank input loop again
		input := scanner.Text()
		if input == "" {
			continue
		}
		cleanedInput := cleanInput(input)

		// Try to match command and call it
		err := command.ExecuteCommand(supportedCommands, cleanedInput, config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	stringLower := strings.ToLower(text)
	return strings.Fields(stringLower)
}
