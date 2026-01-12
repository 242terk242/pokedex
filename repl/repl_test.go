package repl_test

import (
	"testing"

	"github.com/mik242/pokedex/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := repl.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input %q, expected length %d but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input %q, at index %d, expected %q but got %q", c.input, i, expectedWord, word)
			}
		}
	}
}
