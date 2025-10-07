package main

import (
	"fmt"
	"os"
)

// commandfunctions
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
