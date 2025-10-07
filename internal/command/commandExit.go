package command

import (
	"fmt"
	"os"
)

// commandfunctions
func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
