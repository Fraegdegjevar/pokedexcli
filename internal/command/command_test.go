package command

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Fraegdegjevar/pokedexcli/internal/pokeapi"
)

func TestCommandHelp(t *testing.T) {
	//Store os.Stdout so we can replace it after
	// we switch it to a new buffer for this test,
	// as otherwise subsequent tests (which run in a
	//single shared process) could be affected by
	// this mutation to global state.

	// We need to set Stdout to an os.Pipe
	// so we can inspect output.
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	// very important global state restore for other tests
	defer func() { os.Stdout = old }()

	err := commandHelp(&pokeapi.Config{})
	if err != nil {
		t.Fatalf("Error with commandHelp: %v", err)
	}

	// Stop capturing
	w.Close()

	//Read what commandHelp wrote to Os.Stdout by copying
	//into a buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()

	output := buf.String()
	// Check string output has properly formatted welcome
	// message:
	if !strings.Contains(output,
		"\nWelcome to the Pokedex!\nUsage:\n\n") {
		t.Errorf("Expecting welcome text in output but got: %v", output)
	}

	// Check output string has name/description of each
	// function in registry
	if !strings.Contains(output,
		"exit: Exit the Pokedex\nhelp: Displays a help message\nmap: Displays the names of the next 20 location areas in the Pokemon world.\nmapb: Displays the names of the previous 20 location areas in the Pokemon world.") {
		t.Errorf("Expecting command help text but got: %v", output)
	}
}

func TestCommandExit(t *testing.T) {
	// This test function checks if os.Exit(0) is
	// called by spawning a child process that
	// calls this function.

	//Function will return if it determines it is in
	// a child process spawned by this function
	// (and not the parent process)
	if os.Getenv("EXIT_TEST") == "1" {
		commandExit(&pokeapi.Config{})
		return
	}

	//Now launch a new instance of the current test
	// i.e a child process
	// os.Args[0] is the compiled test executable built by
	// Go when testing with go test
	cmd := exec.Command(os.Args[0], "-test.run=TestCommandExit")

	//Mark as child process by setting the env (used above)
	cmd.Env = append(os.Environ(), "EXIT_TEST=1")

	// Run child and get result
	err := cmd.Run()

	//Run() returns nil if command exits with status code 0,
	// otherwise returns an error 'exit status n' for n != 0
	// of type *exec.ExitError
	if err != nil {
		t.Fatalf("Expected commandExit to exit with status code 0, instead got error: %v", err)
	}
}

func TestCommandMap(t *testing.T) {
	//We need to temporarily replace the GetLocationAreas function to return the mock
	// LocationAreaResponse object - ultimately we only want to test commandMap's behaviour

}
