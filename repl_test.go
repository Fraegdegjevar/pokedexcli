package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	// Create a slice of test case structs
	cases := []struct {
		input    string
		expected []string
	}{ //slice begin
		//first struct in slice (first case)
		{
			input:    "  hello 	world ",
			expected: []string{"hello", "world"},
		},
		//second struct in slice (second case)
		{
			input:    "hello jimbob certainly finisher ",
			expected: []string{"hello", "jimbob", "certainly", "finisher"},
		},
		//etc...
		{
			input:    "aRe wE striPPING CASe	",
			expected: []string{"are", "we", "stripping", "case"},
		},
	}

	// Loop over test cases in the cases slice and run tests:
	for _, c := range cases {
		actual := cleanInput(c.input)

		// check the length of the actual slive vs the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Expected string slice length: %v, Actual length: %v ", len(c.expected), len(actual))
		}
		//Check the words in the actual slice match the words in expected
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected and actual word %d don't match: Expected: %v, Actual %v", i, expectedWord, word)
			}
		}
	}

}
